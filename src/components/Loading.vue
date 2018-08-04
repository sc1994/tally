<template>
  <div class="loader13 loader" v-if="show">
    <div class="loader-13">
      <div></div>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    show: {
      type: Boolean,
      required: true
    }
  }
}
</script>
<style>
.loader {
  width: 100%;
  height: 100%;
  background-color: rgba(120, 144, 156, 0.5);
  float: left;
  top: 0px;
  left: 0px;
  display: -webkit-flex;
  justify-content: center;
  display: flex;
  align-items: center;
  z-index: 20181219;
  position: fixed;
}
</style>

