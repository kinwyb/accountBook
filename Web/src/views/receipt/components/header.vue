<template>
      <el-row>
          <span class="search-col search-col--min">
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
          </span>
          <span class="search-col search-col--min">
            <span>银行:</span>
            <el-cascader placeholder="银行" :clearable=true @change="bankIDChange" :options="bankOptions" filterable />
          </span>
          <span class="search-col search-col--min">
            <span>店铺:</span>
            <el-cascader placeholder="店铺" :options="shopOptions" :clearable=true @change="setShopID" />
          </span>
          <span class="search-col search-col--min" >
            <span>类型:</span>
            <el-cascader placeholder="类型" :options="tpOptions" :clearable=true @change="setReceiptTypeID" />
          </span>
          <span>
            <el-button type="success" @click="showAdd()" style="margin-left:20px" >新增</el-button>
          </span>
      </el-row>

</template>

<style lang="less">
.search-col{
  float: left;
  margin-left: 10px;
}
</style>

<script>
import bankAPI from '@/api/bank'
import receiptTypeAPI from '@/api/receiptType'

export default {
  props: ['setDateRange', 'setBankID', 'setReceiptTypeID', 'setShopID', 'showAdd'],
  created () {
    this.loadBank()
    this.loadShop()
    this.loadReceiptType()
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
    loadReceiptType () {
      receiptTypeAPI.ListByLevel({
        level: 2
      }).then(res => {
        if (res.code === 0) {
          for (var v in res.data) {
            this.tpOptions.push({
              value: res.data[v].Name,
              label: res.data[v].Name
            })
          }
        } else {
          console.log(res.errmsg)
        }
      })
    },
    // 加载店铺类型
    loadShop () {
      receiptTypeAPI.ListByLevel({
        level: 0
      }).then(res => {
        if (res.code === 0) {
          for (var v in res.data) {
            this.shopOptions.push({
              value: res.data[v].Id,
              label: res.data[v].Name
            })
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
      shopOptions: [] // 店铺类型
    }
  }
}
</script>
