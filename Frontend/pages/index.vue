<template>
  <div>
    <v-container>
      <b><h1 class="py-5">หน้าแรก</h1></b>
      <v-row>
        <v-col>
          <div
            class="background drop-shadow2 styleblox"
            style="
            background-color: #73C088
              height: 100%;
              color: #fff;
              align-items: center;
            text-align: center;
            "
          >
            <b style="font-size: 55px">{{ lengthdata }}</b>
            <b style="font-size: 18px">กำลังเปิด</b>
          </div>
        </v-col>
        <v-col>
          <item-menu
            icon="mdi-power-standby"
            title="สถานะการใช้งาน"
            to="onoff"
          ></item-menu>
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <item-menu
            icon="mdi-home-lightning-bolt-outline"
            title="บันทึกหน่วยไฟฟ้า"
            to="saveunit"
          ></item-menu>
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <item-menu
            icon="mdi-network-strength-4-cog"
            title="ข้อมูลการตั้งค่า"
            to="setting"
          ></item-menu>
        </v-col>

        <v-col>
          <item-menu
            icon="mdi-home-lightning-bolt-outline"
            title="ความปลอดภัย"
            to="security"
          >
          </item-menu>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import ItemMenu from "~/components/ItemMenu.vue";
export default {
  components: { ItemMenu },
  data() {
    return {
      lengthdata: null,
    };
  },
  mounted() {
    this.$axios
      .get(`${process.env.BASE_URL}/channels/all`)
      .then((response) => {
        this.lengthdata = response.data.data.filter(
          (e) => e.status == true
        ).length;
      })
      .catch((err) => {
        console.log(err);
      });
  },
};
</script>

<style>
</style>
