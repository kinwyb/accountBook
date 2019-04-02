<template>
    <div class="block" style="float:right">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        @prev-click="handleCurrentChange"
        @next-click="handleCurrentChange"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="10"
        layout="total, prev, pager, next,sizes, jumper"
        :total="total">
      </el-pagination>
    </div>
</template>
<script>
export default {
  props: ['handleSizeChange', 'handleCurrentChange', 'total']
}
</script>
