<template>
  <el-container>
    <el-header class="flex" height="auto" >
      <headerView :setDateRange="setDateRange" :setBankID="setBankID" :setReceiptTypeID="setReceiptTypeID" :setShopID="setShopID"  :showAdd="showAdd"/>
    </el-header>
    <el-main>
      <tableView :tableData="tableData" :tableRowClassName="tableRowClassName" />
    </el-main>
    <el-footer>
      <pageView :handleSizeChange="handleSizeChange" :handleCurrentChange="handleCurrentChange" :total="total" :page="page"/>
    </el-footer>

    <!-- Form -->
    <el-dialog title="新增单据" :visible.sync="dialogFormVisible">
      <!-- <el-form :model="form" label-width="120px" size="mini" >
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
      </el-form> -->
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <!-- <el-button type="primary" @click="formSubmit">确 定</el-button> -->
      </div>
    </el-dialog>

  </el-container>
</template>

<script>
import tableView from './components/table'
import headerView from './components/header'
import pageView from './components/page'
import DateUtil from '@/utils/date'
import receiptAPI from '@/api/receipt'

export default {
  components: {
    headerView,
    tableView,
    pageView
  },
  created () {
    this.loadData()
  },
  methods: {
    showAdd (val) {
      this.dialogFormVisible = true
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
    setReceiptTypeID (val) {
      if (val.length < 2) {
        val = val[0]
      } else {
        val = val[1]
      }
      this.receiptType = val
      this.page = 1
      this.total = 0
      this.loadData() // 重新加载数据
    },
    setShopID (val) {
      if (val.length < 2) {
        val = val[0]
      } else {
        val = val[1]
      }
      if (val === '') {
        this.shopID = 0
      } else {
        this.shopID = val
      }
      this.page = 1
      this.total = 0
      this.loadData() // 重新加载数据
    },
    setBankID (val) {
      if (val === '') {
        this.bankID = 0
      } else {
        this.bankID = val
      }
      this.page = 1
      this.total = 0
      this.loadData() // 重新加载数据
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
      if (row.amount >= 0) {
        return 'input-row'
      } else {
        return 'out-row'
      }
    },
    loadData () {
      var path = this.$route.path
      var listType = Number(path.replace('/receipt/list/', ''))
      if (this.tpValue !== listType) {
        this.page = 1
      }
      this.tpValue = listType
      var reqData = {
        'StartTime': this.startDate,
        'EndTime': this.endDate,
        'BankID': Number(this.bankID),
        'ShopID': Number(this.shopID),
        'ReceiptType': this.receiptType,
        'BillType': Number(listType)
      }
      var params = {
        'page': this.page,
        'pageSize': this.pageSize
      }
      receiptAPI.List(reqData, params).then(res => {
        if (res.code === 0) {
          this.total = res.page.total
          this.tableData = []
          for (var v in res.data.Data) {
            var item = res.data.Data[v]
            this.tableData.push({
              billNo: item.Id,
              tp: item.Type,
              amount: item.Money,
              amountType: item.MoneyType,
              userName: item.Operator,
              date: item.Createtime,
              bank: item.Bank,
              desc: item.Description,
              shop: item.Shop
            })
          }
        } else {
          console.log(res.errmsg)
        }
      })
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
      pageSize: 50,
      total: 0,
      tableData: [],
      startDate: '',
      endDate: '',
      bankID: '',
      receiptType: '',
      shopID: '',
      tpValue: 0,
      dialogFormVisible: false
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
