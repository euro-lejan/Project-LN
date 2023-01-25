<template>
  <v-container>
    <div class="wrapper_page">
      <h1>คาวมปลอดภัย</h1>
      <v-row class="warpper-input mt-5 py-10 drop-shadow2 card">
        <span>กำหนดวัตถ์</span>
        <v-text-field
          :disabled="!edite"
          class="mt-5"
          type="number"
          style="font-size: 20px"
          label="ค่าวัตถ์"
          v-model="item.volt"
          suffix="Volt"
        ></v-text-field>
        <v-col
          class="warpper-input d-flex"
          style="padding: 0px 0px 20px 0px; justify-content: space-between"
        >
          <span style="font-size: 20px">สถานะ</span>
          <v-switch
            :disabled="!edite"
            color="success"
            v-model="item.status"
            inset
          ></v-switch>
          <!--v-switch ประกาศ@change="onChange(ตัวแปรที่จะเรียกใช้ในฟังก์ชัน  methods onchang) -->
        </v-col>
        <v-btn
          v-if="!edite"
          block
          elevation="2"
          color="#f9ae00"
          style="color: #fff; border-radius: 15px"
          @click="(edite = true), (backup = JSON.parse(JSON.stringify(item)))"
          x-large
          >แก้ไข</v-btn
        >
        <v-row v-else>
          <v-col>
            <v-btn
              block
              elevation="2"
              color="#526FFF"
              style="color: #fff; border-radius: 15px"
              type="submit"
              x-large
              @click="onsubmit()"
              >บันทึก</v-btn
            >
          </v-col>
          <v-col>
            <v-btn
              block
              elevation="2"
              color="red"
              style="color: #fff; border-radius: 15px"
              type="submit"
              x-large
              @click="(edite = false), (item = backup)"
              >ยกเลิก</v-btn
            >
          </v-col>
        </v-row>
        <!--v-switch ประกาศ@change="onChange(ตัวแปรที่จะเรียกใช้ในฟังก์ชัน  methods onchang) -->
      </v-row>
    </div>
  </v-container>
</template>
  
  <script>
import axios from "axios";
export default {
  data() {
    return {
      edite: false,
      backup: null,
      item: {
        volt: null,
        status: false,
      },
    };
  },
  mounted() {
    this.$axios
      .get(`${process.env.BASE_URL}/savety/all`)
      .then((response) => {
        this.item = response.data.data;
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    onsubmit() {
      this.$axios
        .put(`${process.env.BASE_URL}/savety/update`, {
          ...this.item,
          volt: parseInt(this.item.volt),
        })
        .then((response) => {
          this.edite = false;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
  
  <style lang="scss" scoped>
.wrapper_page {
  flex-direction: column;
  align-content: center;

  .warpper-input {
    align-items: center;
    background: #fff;
    border-radius: 15px;
    font-size: 25px;
    text-align: center;
    b {
      margin: 0;
      // color: #fff;
    }
  }
}
</style>