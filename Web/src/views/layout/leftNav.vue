<template>
<el-aside width="200px">
  <el-menu router id="main_menu" :default-openeds="['1','3']">
      <el-submenu index="1">
        <template slot="title"><i class="el-icon-message"></i>收支单据</template>
          <el-menu-item index="/receipt/list/0">所有单据</el-menu-item>
          <el-menu-item index="/receipt/list/1">收入单据</el-menu-item>
          <el-menu-item index="/receipt/list/2">支出单据</el-menu-item>
      </el-submenu>
      <!-- <el-submenu index="2">
        <template slot="title"><i class="el-icon-message"></i>销售单</template>
          <el-menu-item index="/saleBill">所有单据</el-menu-item>
          <el-menu-item index="/receipt/list/1">收入单据</el-menu-item>
          <el-menu-item index="/receipt/list/2">支出单据</el-menu-item>
      </el-submenu> -->
      <el-submenu index="3">
        <template slot="title"><i class="el-icon-menu"></i>系统设置</template>
        <el-menu-item index="/system/bank">银行</el-menu-item>
        <el-menu-item index="/system/receiptType">收支类型</el-menu-item>
      </el-submenu>
    </el-menu>
</el-aside>
</template>

<style lang="scss">
#main_menu {
  height: 98vh;
  background-color: white
}
</style>

<script>
export default {
  data () {
    return {
    }
  }
}
</script>
