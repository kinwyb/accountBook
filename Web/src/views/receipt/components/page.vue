<template>
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
</template>
<script>
export default {
  props: ['handleSizeChange', 'handleCurrentChange', 'total', 'page']
}
</script>
