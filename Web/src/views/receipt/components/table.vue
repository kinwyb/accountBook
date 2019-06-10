<template>
  <div>
    <el-table
      :data="tableData"
      style="width: 100%"
      height="70vh"
      :row-class-name="tableRowClassName"
      :summary-method="getSummaries"
    >
      <el-table-column prop="billNo" label="单据号" width="150"/>
      <el-table-column prop="shop" label="店铺" width="150"/>
      <el-table-column prop="tp" label="类型" width="150"/>
      <el-table-column prop="amount" label="金额" width="160"/>
      <el-table-column prop="amountType" label="货币" width="120"/>
      <el-table-column prop="date" label="日期" width="180"/>
      <el-table-column prop="bank" label="银行" width="150"/>
      <el-table-column prop="desc" label="备注" width="250"/>
    </el-table>
    <el-table
      :data="countData"
      style="width: 100%"
      :header-row-class-name='"countRow"'
      :row-class-name='"countRow"'
    >
      <el-table-column prop="MoneyType" label="货币"/>
      <el-table-column prop="AllIn" label="总收入"/>
      <el-table-column prop="AllOut" label="总支出"/>
      <el-table-column prop="StartMoney" label="期初余额"/>
      <el-table-column prop="EndMoney" label="余额"/>
    </el-table>
  </div>
</template>

<style scoped>
.el-table >>> .countRow {
  background-color: #f0f9eb;
}
</style>

<script>
export default {
  props: ['tableData', 'countData', 'tableRowClassName'],
  methods: {
    getSummaries (param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '总价'
          return sums
        }
        const values = data.map(item => Number(item[column.property]))
        if (!values.every(value => isNaN(value))) {
          sums[index] = values.reduce((prev, curr) => {
            const value = Number(curr)
            if (!isNaN(value)) {
              return prev + curr
            } else {
              return prev
            }
          }, 0)
          sums[index] = sums[index].toFixed(3)
          sums[index] += ' 元'
        } else {
          sums[index] = ''
        }
      })
      return sums
    }
  }
}
</script>
