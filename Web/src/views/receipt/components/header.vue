<template>
    <el-header>
      <el-date-picker
        v-model="dateRangeValue"
        type="daterange"
        align="right"
        unlink-panels
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        :picker-options="pickerOptions" />
      <el-button @click="resetArgs">重置时间</el-button>
      <span>银行:</span>
       <el-cascader placeholder="银行" :clearable=true @change="bankIDChange" :options="bankOptions" filterable />
      <span>类型:</span>
      <el-cascader placeholder="类型" :options="tpOptions" @change="setReceiptTypeID" @active-item-change="loadReceiptType" filterable />
    </el-header>
</template>

<script>
import bankAPI from '@/api/bank'
import receiptTypeAPI from '@/api/receiptType'

export default {
  props: ['setDateRange', 'setBankID', 'setReceiptTypeID'],
  created () {
    this.loadBank()
    this.loadReceiptType(0)
  },
  methods: {
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
    // 加载收支类型
    loadReceiptType (val) {
      var tpMap = Array(0)
      for (var v in this.tpOptions) {
        tpMap[this.tpOptions[v].value] = this.tpOptions[v]
      }
      this.parentID = val.toString()
      receiptTypeAPI.List({
        parentID: val.toString()
      }).then(res => {
        if (val === 0) {
          this.tpOptions = [{
            value: 0,
            label: '不限'
          }]
        } else {
          tpMap[val].children = []
        }
        if (res.code === 0) {
          if (val !== 0 && res.data === null) {
            tpMap[this.parentID].children = null
          }
          for (var v in res.data) {
            if (val === 0) {
              this.tpOptions.push({
                value: res.data[v].Id,
                label: res.data[v].Name,
                children: []
              })
            } else {
              tpMap[this.parentID].children.push({
                value: res.data[v].Id,
                label: res.data[v].Name
              })
            }
          }
        } else {
          console.log(res.errmsg)
        }
      })
    },
    resetArgs () {
      this.dateRangeValue = ''
    },
    bankIDChange (index) {
      this.setBankID(index.toString())
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
      pickerOptions: {
        shortcuts: [{
          text: '今天',
          onClick (picker) {
            const date = new Date()
            picker.$emit('pick', [date, date])
          }
        }, {
          text: '昨天',
          onClick (picker) {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24)
            picker.$emit('pick', [date, date])
          }
        }, {
          text: '本月',
          onClick (picker) {
            const date = new Date()
            const start = new Date()
            start.setDate(1)
            picker.$emit('pick', [start, date])
          }
        }, {
          text: '最近一周',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近一个月',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近三个月',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
            picker.$emit('pick', [start, end])
          }
        }]
      },
      tpOptions: [], // 收支单据类型
      bankOptions: [], // 银行列表数据
      parentID: 0
    }
  }
}
</script>
