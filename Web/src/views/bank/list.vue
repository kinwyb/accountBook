<template>
    <el-container>
    <el-header class="flex" height="auto" >
        <span class="search-col search-col--min">
            <el-date-picker
                    v-model="dateRangeValue"
                    type="daterange"
                    align="right"
                    unlink-panels
                    range-separator="至"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期" />
        </span>
        <el-button type="primary" round @click="loadData" >刷新</el-button>
        <el-button type="success" @click="dialogFormVisible = true" >新增</el-button>
    </el-header>
    <el-main>
      <el-table :data="tableData" style="width: 100%" height="70vh">
      <el-table-column prop="BankName" label="银行名称" width="250" />
      <el-table-column prop="SCNY" label="期初人民币" width="120"/>
      <el-table-column prop="ECNY" label="期末人民币" width="120"/>
      <el-table-column prop="SUSD" label="期初美金" width="100"/>
      <el-table-column prop="EUSD" label="期末美金" width="100"/>
      <el-table-column prop="BankAccount" label="银行账户" width="180"/>
      <el-table-column prop="Contacts" label="联系人" width="150" />
      <el-table-column prop="ContactPhone" label="联系电话" width="120"/>
    </el-table>
    </el-main>

    <!-- Form -->
    <el-dialog title="新增银行" :visible.sync="dialogFormVisible">
      <el-form :model="form" label-width="120px" size="mini" >
        <el-form-item label="名称">
          <el-input v-model="form.BankName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="账户">
          <el-input v-model="form.BankAccount" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="form.BankPeople" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="form.BankPhone" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="期初人民币">
          <el-input v-model="form.BankMoney" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="期初美金">
          <el-input v-model="form.BankMoneyUsa" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="formSubmit">确 定</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import bankAPI from '@/api/bank'
import DateUtil from '@/utils/date'

export default {
  created () {
    this.emptyForm = JSON.parse(JSON.stringify(this.form))
    this.loadData()
  },
  methods: {
    formSubmit () {
      this.form.BankMoney = parseFloat(this.form.BankMoney)
      this.form.BankMoneyUsa = parseFloat(this.form.BankMoneyUsa)
      bankAPI.BankAdd(this.form).then(res => {
        if (res.code === 0) {
          this.$notify({
            title: res.data,
            type: 'success'
          })
          console.log(this.emptyForm)
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
      this.dialogFormVisible = false
    },
    setDateRange (val) {
      this.startDate = ''
      this.endDate = ''
      if (val.length > 0) {
        this.startDate = DateUtil.DateFormat(val[0], 'yyyy-MM-dd')
      }
      if (val.length > 1) {
        this.endDate = DateUtil.DateFormat(val[1], 'yyyy-MM-dd')
      }
      this.loadData()
    },
    loadData () {
      console.log('加载数据...')
      this.tableData = []
      bankAPI.BankListCompate().then(res => {
        if (res.code === 0) {
          this.tableData = res.data
        } else {
          console.log('数据加载失败=>', res)
        }
        this.tableData.forEach(val => {
          val.SCNY = val.SCNY === 0 ? '' : val.SCNY.toFixed(3)
          val.SUSD = val.SUSD === 0 ? '' : val.SUSD.toFixed(3)
          val.ECNY = val.ECNY === 0 ? '' : val.ECNY.toFixed(3)
          val.EUSD = val.EUSD === 0 ? '' : val.EUSD.toFixed(3)
        })
      })
    }
  },
  watch: {
    'dateRangeValue' () {
      this.setDateRange(this.dateRangeValue)
    }
  },
  data () {
    return {
      dateRangeValue: '',
      tableData: [],
      startDate: '',
      endDate: '',
      form: {
        BankName: '',
        BankAccount: '',
        BankPhone: '',
        BankPeople: '',
        BankMoney: 0.0,
        BankMoneyUsa: 0.0
      },
      emptyForm: {},
      dialogFormVisible: false
    }
  }
}
</script>
