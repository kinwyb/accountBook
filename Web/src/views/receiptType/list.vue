<template>
  <el-container>
    <el-main>
       <el-tree
      :data="tableData"
      node-key="ID"
      accordion
      :highlight-current="true"
      :expand-on-click-node="true"
      :render-content="renderContent">
      </el-tree>
    </el-main>
    <!-- Form -->
    <el-dialog :title="addTitle" :visible.sync="dialogFormVisible">
      <el-form :model="form" label-width="120px" size="mini" >
        <el-form-item label="名称">
          <el-input v-model="form.Name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="formSubmit">确 定</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<style>
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }
</style>

<script>
import ReceiptTypeAPI from '@/api/receiptType'

export default {
  created () {
    this.emptyForm = JSON.parse(JSON.stringify(this.form))
    this.loadData()
  },
  methods: {
    renderContent (h, { node, data, store }) {
      var add = <el-button size="mini" type="text" on-click={ () => this.append(data) }>增加下一级</el-button>
      if (node.level !== 1) {
        add = ''
      }
      return (
        <span class="custom-tree-node">
          <span>{node.data.Name}</span>
          <span>
            {add}
          </span>
        </span>)
    },
    formSubmit () {
      ReceiptTypeAPI.Add(this.form).then(res => {
        if (res.code === 0) {
          this.$notify({
            title: res.data,
            type: 'success'
          })
          console.log(this.emptyForm)
          // const newChild = { id: id++, label: 'testtest', children: [] }
          // if (!data.children) {
          //   this.$set(data, 'children', [])
          // }
          // data.children.push(newChild)
          this.form = this.emptyForm
        } else {
          this.$alert(res.errmsg, '操作失败', {
            confirmButtonText: '确定',
            callback: action => {
              this.$notify({
                type: 'error',
                message: `操作失败: ${res.errmsg}`
              })
            }
          })
        }
      })
      this.form = this.emptyForm
      this.dialogFormVisible = false
    },
    loadData () {
      console.log('加载数据...')
      this.tableData = []
      ReceiptTypeAPI.Tree().then(res => {
        if (res.code === 0) {
          this.tableData = res.data
          console.log(res.data)
        } else {
          this.$notify({
            type: 'error',
            message: `数据加载失败:${res.errmsg}`
          })
        }
      })
    },
    append (data) {
      this.dialogFormVisible = true
      this.form.ParentId = data.ID
      this.addTitle = '新增 [' + data.Name + '] 下级'
    }
  },
  data () {
    return {
      addTitle: '',
      tableData: [],
      form: {
        ParentId: 0,
        Name: ''
      },
      emptyForm: {},
      dialogFormVisible: false
    }
  }
}
</script>
