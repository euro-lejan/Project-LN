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
            <b style="font-size: 18px">อัพเดตเมื่อ : {{ time_update }}</b>
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
      time_update: "กำลังโหลด",
    };
  },
  mounted() {
    // สร้างฟังก์ชันสำหรับดึงข้อมูลจาก API /channels/all
    const fetchData = () => {
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
    };

    // สร้างฟังก์ชันสำหรับดึงข้อมูลจาก API /nodemcu/status
    const fetchStatus = () => {
      this.$axios
        .get(`${process.env.BASE_URL}/nodemcu/status`)
        .then((response) => {
          const millis = response.data.data[0].millis; // รับค่า millis จาก API แทน timestamp

          const currentMillis = new Date().getTime(); // เวลาปัจจุบันในรูปแบบ millis
          // millisDifference > 15000
          const millisDifference = currentMillis - parseInt(millis); // คำนวณระยะห่างระหว่างเวลา

          if (millisDifference > 15000) {
            if (this.mcu_status) {
              alert("ขาดการเชื่อมต่อ กรุณาตรวจสอบ โมดูล");
            }
            this.mcu_status = false; // ถ้าระยะห่างเกิน 15 วินาที
          } else {
            if (!this.mcu_status) {
              alert("กลับมาเชื่อมต่ออีกครั้ง");
            }
            this.mcu_status = true; // ถ้าระยะห่างไม่เกิน 15 วินาที
          }

          const currentTime = new Date(currentMillis);
          const year = currentTime.getFullYear();
          const month = (currentTime.getMonth() + 1)
            .toString()
            .padStart(2, "0");
          const day = currentTime.getDate().toString().padStart(2, "0");
          const hour = currentTime.getHours().toString().padStart(2, "0");
          const minute = currentTime.getMinutes().toString().padStart(2, "0");
          const second = currentTime.getSeconds().toString().padStart(2, "0");

          this.time_update = `${day}/${month}/${year} ${hour}:${minute}:${second}`;
        })
        .catch((err) => {
          console.log(err);
        });
    };
    fetchData();
    setInterval(fetchStatus, 10000);
  },
};
</script>


<style>
</style>
