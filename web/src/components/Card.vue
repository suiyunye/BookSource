<template>
  <el-card class="box-card">
    <div class="info">
      <div class="name">{{ repoName }}</div>
      <div class="size">书源数目：{{ size }}</div>
    </div>
    <div class="url">{{ url }}</div>
    <div class="detail">{{ summary }}</div>
    <div class="copy ripple" @click="copy">复制地址</div>
  </el-card>
</template>

<script>
export default {
  props: ["url", "repoName", "summary", "size"],
  methods: {
    copy() {
      const that = this;
      this.$copyText(this.url).then(
        () => {
          that.$message({ message: "仓库地址复制成功", type: "success" });
        },
        () => {
          that.$message.error("复制失败了");
        }
      );
    }
  }
};
</script>

<style scoped>
.box-card {
  width: calc(100vw - 36px);
  margin-top: 28px;
  display: flex;
  flex-direction: column;
  font-size: 14px;
  line-height: 1.5;
}

.info {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  color: #333;
  /* margin-top: 9px; */
}

.url {
  margin: 18px 0;
  word-break: break-all;
  color: #aaaaaa;
}

.detail {
  margin: 20px 0 32px 0;
  white-space: pre-wrap;
  line-height: 1.7;
  color: #aaaaaa;
}

.copy {
  height: 48px;
  line-height: 48px;
  border-top: #f3f3f3 1px solid;
  text-align: center;
  margin-bottom: -20px;
  width: calc(100vw - 36px);
  margin-left: -20px;
  cursor: pointer;
  user-select: none;
}

.ripple {
  position: relative;
  overflow: hidden;
}

.ripple:after {
  content: "";
  display: block;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  pointer-events: none;
  background-image: radial-gradient(circle, #666 10%, transparent 10.01%);
  background-repeat: no-repeat;
  background-position: 50%;
  transform: scale(10, 10);
  opacity: 0;
  transition: transform 0.3s, opacity 0.5s;
}

.ripple:active:after {
  transform: scale(0, 0);
  opacity: 0.3;
  transition: 0s;
}
</style>
