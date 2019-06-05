<template>
    <div>
        <el-form :model="form" label-width="120px" size="mini" :inline="true" >
            <el-row>
                <el-col :span="12">
                    <el-form-item label="单据号" label-width="auto">
                        <el-input v-model="form.id" autocomplete="off" class="inputWidth" :disabled="true" />
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="金额" label-width="auto" >
                        <el-input v-model="form.money" autocomplete="off" class="inputWidth" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="类型" class="labelWidth" label-width="auto">
                        <el-cascader
                        placeholder="类型"
                        :options="tpOptions"
                        :props="{ checkStrictly: true }"
                        @change="receiptTypeChange"
                        clearable></el-cascader>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="银行" label-width="auto" >
                        <el-cascader placeholder="银行" :clearable=true @change="bankIDChange" :options="bankOptions" filterable />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="时间" class="labelWidth" label-width="auto">
                        <el-date-picker
                        v-model="form.datetime"
                        type="datetime"
                        value-format="yyyy-MM-dd HH:mm:ss">
                        </el-date-picker>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="货币" label-width="auto" >
                        <el-radio-group v-model="form.moneyType">
                            <el-radio label="2">人民币</el-radio>
                            <el-radio label="1">美金</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <el-form-item label="备注" class="labelWidth" label-width="auto" >
                        <el-input type="textarea" v-model="form.content" autocomplete="off" style="width:340%"/>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <!-- <el-button @click="dialogFormVisible = false">上一条</el-button>
            <el-button @click="dialogFormVisible = false">下一条</el-button>
            <el-button @click="dialogFormVisible = false">新增</el-button> -->
            <el-button type="primary" @click="formSubmit" style="float:right">提交</el-button>
        </div>
    </div>
</template>

<style>
.labelWidth {
    margin-left: 15px;
}
.inputWidth {
    width: 150%;
}
</style>

<script>
import receiptAPI from '@/api/receipt'
import bankAPI from '@/api/bank'
import receiptTypeAPI from '@/api/receiptType'
import DateUtil from '@/utils/date'

export default {
  props: ['hideAdd'],
  created () {
    this.getLastNo()
    this.loadBank()
    this.loadReceiptType()
  },
  methods: {
    getLastNo () {
      receiptAPI.LastNo().then((res) => {
        this.lastNo = res.data
        this.form.id = res.data
      })
    },
    // 加载银行信息
    loadBank () {
      bankAPI.BankList().then(res => {
        if (res.code !== 0) {
          console.log(res.errmsg)
        } else {
          for (var v in res.data) {
            this.bankOptions.push({
              value: res.data[v].Id,
              label: res.data[v].BankName
            })
          }
        }
      })
    },
    bankIDChange (index) {
      this.form.bank = this.bankOptions[index].value
    },
    // 加载收支类型
    loadReceiptType () {
      receiptTypeAPI.Tree().then(res => {
        if (res.code === 0) {
          for (var v in res.data) {
            var o = {
              value: res.data[v].ID,
              label: res.data[v].Name,
              children: []
            }
            for (var v1 in res.data[v].children) {
              o.children.push({
                value: res.data[v].children[v1].ID,
                label: res.data[v].children[v1].Name
              })
            }
            this.tpOptions.push(o)
          }
        } else {
          console.log(res.errmsg)
        }
      })
    },
    receiptTypeChange (value) {
      this.form.tp = value[1]
    },
    // 提交
    formSubmit () {
      if (this.form.datetime instanceof Date) {
        this.form.datetime = DateUtil.DateFormat(
          this.form.datetime, 'yyyy-MM-dd hh:mm:ss')
      }
      receiptAPI.Add({
        Money: Number.parseFloat(this.form.money),
        BankId: this.form.bank,
        Description: this.form.content,
        Createtime: this.form.datetime,
        Type: this.form.tp,
        MoneyType: Number.parseInt(this.form.moneyType)
      }).then((res) => {
        if (res.code === 0) {
          this.$notify({
            title: res.data,
            type: 'success'
          })
          this.form = this.emptyForm
          this.hideAdd()
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
    }
  },
  data () {
    return {
      lastNo: '',
      bankOptions: [], // 银行列表数据
      tpOptions: [], // 收支单据类型
      form: {
        id: '',
        money: '',
        tp: '',
        bank: '',
        content: '',
        datetime: new Date(),
        moneyType: '2'
      },
      emptyForm: {
        id: '',
        money: '',
        tp: '',
        bank: '',
        content: '',
        datetime: new Date(),
        moneyType: '2'
      }
    }
  }
}
</script>
