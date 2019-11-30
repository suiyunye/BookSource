const axios = require("axios");
const fs = require("fs");
const path = require("path");
const dayjs = require("dayjs");

const hm =
  "https://raw.githubusercontent.com/chimisgo/BookSource/master/src/hm.json";

async function get(url) {
  try {
    let res = await axios.get(url);
    res = res.data;
    return new Promise(resolve => {
      if (res.code === 0) {
        resolve(res);
      } else {
        resolve(res);
      }
    });
  } catch (err) {
    console.log(err);
  }
}

async function getRepos() {
  var res = await get(hm);
  var repos = res.list;
  var datas = [];
  var repoSize = repos.length;
  repos.map(async r => {
    var url = r.url;
    var repo = await get(url);
    var repoName = repo.name;
    var summary = repo.summary;
    var size = repo.list.length;
    var data = {
      url: url,
      repoName: repoName,
      summary: summary,
      size: size
    };
    datas.push(data);
    if (datas.length == repoSize) {
      var time = dayjs().format("YYYY/MM/DD HH:mm:ss");
      var json = {
        time: time,
        repos: datas
      };
      fs.writeFileSync(
        path.resolve(__dirname, "../src/assets/repo.json"),
        JSON.stringify(json, null, 2)
      );
      console.log("数据更新成功！");
    }
  });
}

getRepos();
