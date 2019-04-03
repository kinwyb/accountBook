<template>
  <el-container>
    <el-header>
      <headerView :setDateRange="setDateRange" />
    </el-header>
    <el-main>
      <tableView :tableData="tableData" :tableRowClassName="tableRowClassName" />
    </el-main>
    <el-footer>
      <pageView :handleSizeChange="handleSizeChange" :handleCurrentChange="handleCurrentChange" :total="total"/>
    </el-footer>
  </el-container>
</template>

<script>
import tableView from './components/table'
import headerView from './components/header'
import pageView from './components/page'
import DateUtil from '@/utils/date'
import bankAPI from '@/api/bank'

export default {
  components: {
    headerView,
    tableView,
    pageView
  },
  created () {
    this.loadBank()
    this.loadData()
  },
  methods: {
    // 加载银行信息
    loadBank () {
      bankAPI().then(res => {
        console.log(res)
      })
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
      this.page = 1
      this.total = 0
      this.loadData()
    },
    handleSizeChange (val) {
      this.pageSize = val
      this.loadData()
    },
    handleCurrentChange (val) {
      this.page = val
      this.loadData()
    },
    tableRowClassName ({row, rowIndex}) {
      if (rowIndex % 2 === 1) {
        return 'input-row'
      } else if (rowIndex % 3 === 1) {
        return 'out-row'
      }
      return ''
    },
    loadData () {
      var path = this.$route.path
      var listType = Number(path.replace('/receipt/list/', ''))
      console.log(this.startDate + ' - ' + this.endDate)
      console.log('加载数据' + listType + '第' + this.page + '页,' + this.pageSize)
      if (listType === 1) {
        var item = {
          billNo: 'S000000007003',
          tp: '泓旭',
          amount: 20000.8,
          amountType: '人民币',
          userName: '丁丽丽',
          date: '2019-03-28 16:02:52',
          bank: '支付宝',
          desc: '备注呵呵呵'
        }
        this.tableData = Array(121).fill(item)
      } else if (listType === 2) {
        item = {
          billNo: 'S000000008003',
          tp: '物业',
          amount: 109483,
          amountType: '人民币',
          userName: '丁丽丽',
          date: '2019-01-18 10:02:52',
          bank: '老板娘',
          desc: '备注呵呵呵'
        }
        this.tableData = Array(79).fill(item)
      } else {
        item = {
          billNo: 'S000000008803',
          tp: '货款',
          amount: 4839.1,
          amountType: '人民币',
          userName: '丁丽丽',
          date: '2019-02-08 7:02:52',
          bank: '微信',
          desc: '备注呵呵呵'
        }
        this.tableData = Array(200).fill(item)
      }
      this.total = this.tableData.length
    }
  },
  watch: {
    '$route' (to, from) {
      this.loadData(1)
    }
  },
  data () {
    return {
      page: 1,
      pageSize: 10,
      total: 100,
      tableData: [],
      startDate: '',
      endDate: ''
    }
  }
}
</script>
<style>
  .el-table .input-row {
    color: green;
  }
  .el-table .out-row {
    color: red;
  }
</style>
