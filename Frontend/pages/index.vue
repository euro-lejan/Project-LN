<template>
  <div>
    <v-container>
      <b><h1 class="py-5">หน้าแรก</h1></b>
      <v-row>
        <v-col>
          <div
            class="background drop-shadow2 styleblox"
            :style="{
              'background-color': mcu_status ? '#73C088' : '#FF0000',
              height: '100%',
              color: '#fff',
              'align-items': 'center',
              'text-align': 'center',
            }"
          >
            <b style="font-size: 55px">
              <span v-if="mcu_status">สถานะเชื่อมต่ออยู่</span>
              <span v-else>ขาดการเชื่อมต่อ</span>
            </b>
          </div>
        </v-col>
      </v-row>
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
            icon="mdi-timer-cog"
            title="ข้อมูลการตั้งค่า"
            to="setting"
          ></item-menu>
        </v-col>

        <v-col>
          <item-menu
            icon="mdi-meter-electric-outline"
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
      mcu_status: true,
    };
  },
  mounted() {
    this.$axios
      .get(`${process.env.BASE_URL}/channels/all`)
      .then((response) => {
        this.lengthdata = response.data.data.filter(
          (e) => e.status == true
        ).length;

        return this.$axios.get(`${process.env.BASE_URL}/nodemcu/status`); // เปลี่ยน URL ให้เป็น http://localhost:8888/nodemcu/status
      })
      .then((response) => {
        const millis = response.data.millis; // รับค่า millis จาก API แทน timestamp

        const currentMillis = new Date().getTime(); // เวลาปัจจุบันในรูปแบบ millis

        const millisDifference = currentMillis - parseInt(millis); // คำนวณระยะห่างระหว่างเวลา

        if (millisDifference > 15000) {
          this.mcu_status = false; // ถ้าระยะห่างเกิน 15 วินาที
        } else {
          this.mcu_status = true; // ถ้าระยะห่างไม่เกิน 15 วินาที
        }
      })
      .catch((err) => {
        console.log(err);
      });
  },
};
</script>


<style>
</style>
