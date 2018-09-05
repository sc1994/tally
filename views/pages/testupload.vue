<template>
  <div>
    <upload :server="upload.server" :api="upload.api" :filename="upload.filename" :params="upload.params" @success="upsuccess" :crop="true" :width="upload.width" :height="upload.height" :ok="upload.ok" :cancel="upload.cancel" :auto="true" :quality="0.5" :ext="upload.ext">
      <mu-icon value="cloud_upload"></mu-icon>
    </upload>
  </div>
</template>
<script>
import upload from "@/components/uploadavatar";
export default {
  components: { upload },
  data() {
    return {
      upload: {
        server: "http://118.24.27.231:81",
        api: "/uploadfile",
        params: {
          id: "123",
          fileName: "123"
        },
        crop: true,
        width: 500,
        height: 500,
        cancel: "取消",
        ok: "确定",
        filename: "files",
        ext: "jpg,png"
      }
    };
  },
  methods: {
    // 上传成功后触发
    upsuccess(data) {
      debugger;
      console.log("success");
    }
  }
};
</script>
