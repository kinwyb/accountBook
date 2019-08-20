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
    </el-header>
    <el-main>
      <el-table :data="tableData" style="width: 100%" height="70vh">
        <el-table-column prop="InstallTime" label="时间"/>
        <el-table-column prop="User" label="操作人员"/>
        <el-table-column prop="Action" label="操作动作"/>
        <el-table-column prop="Content" label="具体动作"/>
      </el-table>
      <div class="block" style="float:right">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          @prev-click="handleCurrentChange"
          @next-click="handleCurrentChange"
          :page-sizes="[50, 100,150, 200]"
          :page-size='50'
          :current-page="page"
          layout="total, prev, pager, next,sizes, jumper"
          :total="total">
        </el-pagination>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import systemAPI from '@/api/system'
import DateUtil from '@/utils/date'

export default {
  created () {
    this.loadData()
  },
  methods: {
    handleSizeChange (val) {
      this.pageSize = val
      this.loadData()
    },
    handleCurrentChange (val) {
      this.page = val
      this.loadData()
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
      if (this.page < 2) {
        this.tableData = []
      }
      var params = {
        'page': this.page,
        'pageSize': this.pageSize
      }
      systemAPI.LogList({
        'StartTime': this.startDate,
        'EndTime': this.endDate
      }, params).then(res => {
        if (res.code === 0) {
          this.total = res.page.total
          res.data.forEach(val => {
            val.User = '丁丽丽'
            this.tableData.push(val)
          })
        } else {
          console.log('数据加载失败=>', res)
        }
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
      total: 0,
      page: 1,
      pageSize: 50
    }
  }
}
</script>
