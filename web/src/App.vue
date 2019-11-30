<template>
  <div id="app">
    <div class="title">
      书源仓库索引
    </div>
    <div class="update-time">
      {{ updateTime }}
    </div>
    <card v-for="repo in repos" :key="repo.repoName" v-bind="repo"></card>
  </div>
</template>

<script>
import json from "./assets/repo.json";
import card from "./components/Card.vue";
export default {
  name: "app",
  data() {
    return {
      updateTime: "更新于 " + json.time,
      repos: json.repos
    };
  },
  components: {
    card
  }
};
</script>

<style>
#app {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 42px;
  padding-bottom: 42px;
}

.title {
  font-size: 18px;
  font-weight: 400;
  line-height: 1.5;
}

.update-time {
  font-size: 14px;
  color: #c4c4c5;
  line-height: 1.5;
  margin-top: 12px;
}
</style>
