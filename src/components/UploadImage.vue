<template>
  <dialoghead :open.sync="thatOpen" :buttonclick="submit" buttonicon="done" :title="'上传头像'" :buttonshow="true">
    <vuecoreimageupload :url="uploadUrl" :crop="false" @imageuploaded="imageuploaded" text="选择你的头像" crop="true|local " cropRatio="1:1 " inputAccept="image/jpg,image/jpeg,image/png " :max-file-size="5242880 ">
    </vuecoreimageupload>
  </dialoghead>
</template>

<script>
import vuecoreimageupload from "vue-core-image-upload";
import dialoghead from "@/layout/dialog";
import { mapState } from "vuex";

export default {
  components: {
    vuecoreimageupload,
    dialoghead
  },
  props: {
    open: {
      type: Boolean
    }
  },
  data() {
    return {
      thatOpen: false,
      uploadUrl: "http://localhost/uploaduserheadimage/123123123",
      src:
        "http://img1.vued.vanthink.cn/vued0a233185b6027244f9d43e653227439a.png"
    };
  },
  methods: {
    imageuploaded(res) {
      if (res.errcode == 0) {
        this.src = res.data.src;
      }
    },
    submit() {}
  },
  computed: {
    ...mapState(["currentUser"])
  },
  watch: {
    open(val) {
      if (val) {
        this.thatOpen = true;
      }
    },
    thatOpen(val) {
      this.$emit("update:open", false);
    }
  }
};
</script>