<template>
  <div id="wrapper" class="container">
    <div class="back">
      <div class="course">
        <span><!--aktueller Kurs:--> <b>{{ course }}</b></span>
      </div>
      <router-link :to="'/'+encodeURIComponent(account)"><i class="fa fa-chevron-left"></i> Zurück zu meinen Kursen</router-link>
      <i class="fas fa-cog right" @click="settings = !settings"></i>


      <transition name="grow">
        <div v-show="settings" class="container settings">
          
          <div class="row">
            <h3>Einstellungen</h3>
          </div>
          <div class="row">
            <span>Prozent </span>  
              <label class="switch">
                  <input type="checkbox" v-model="showPoints">
                  <span class="slider round"></span>
                </label>
                <span>Punkte</span>
          </div>
          <div class="row">

              <div v-show="!showPoints">
                <label for="Prozent" style="font-weight: normal; display: inline;">mindestens zu erreichende Prozent: </label>
                <input id="Prozent" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="100"  v-model="percent">
              </div>
              <div v-show="showPoints">
                <label for="zuErreichen" style="font-weight: normal; display: inline;">mindestens zu erreichende Punkte: </label>
                <input id="zuErreichen" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="100" v-model="zuErreichen">
                <label for="von" style="font-weight: normal; display: inline;">/</label>
                <input id="von" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="100" v-model="von"><br>
                <small >Achtung: Die Punkte werden für das Backend und die Validierung in Prozent umgerechnet, überprüfen Sie bitte das Ergebniss.</small>
              </div>
            </div>
            <div class="row">
                <label for="vorträge" style="font-weight: normal; display: box;">Wie viele Vorträge müssen erbracht werden: </label>
                <input id="vorträge" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="10" v-model="vorträge">
              </div>
            <div class="row">
              <button class="button" @click="changeP">Speichern</button>
          </div>
        </div>
      </transition>


    </div>
    <table class="tasks">
      <thead>
        <tr>
          <th v-for="key in gridColumns">
            {{ key }}
          </th>
          <th class="small"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(entry, index) in gridData">
          <td >
            {{ entry['name'] }}
          </td>
          <td class="formatted-points">
            <span>{{ entry['points'] }}</span><span>/</span><span>{{ entry['maxPoints'] }}</span>
          </td>
          <td style="text-align: center">
            {{ entry['presentations'] }}
          </td>
          <td class="small">
            <i class="fas fa-edit clickable-icon" @click="edit(entry)" style="margin-right: 5px; display: inline"></i
            ><i class="fas fa-trash-alt clickable-icon" @click="deleteItem(index, entry['id'])" style="display: inline"></i>
          </td>
        </tr>
        <tr>
          <td>
            <b>Gesamt</b>
          </td>
          <td class="formatted-points">
            <b>{{ gesPoints }}</b><span>/</span><b>{{ gesMaxPoints }}</b>
          </td>
          <td style="text-align: center">
            <b>{{ gesPres }}</b><span> /</span><b style="opacity:0.7; margin-left:3px;">{{ pres }}</b>
          </td>
          <td>
            <b style="margin-right:3px;">{{ gesPerc }}</b><span>/</span><b style="opacity:0.7; margin-left:3px;">{{ perc }}%</b>
          </td>
        </tr>
      </tbody>
    </table>
    <fieldset class="row">
      <div class="column full" style="flex-grow: 2">
        <label for="Aufgabe">Aufgabe</label>
        <input id="Aufgabe" type="text" v-model="editData.name" v-on:keyup.enter="setFocus('punkte')">
      </div>
      <div class="column column-17">
        <label for="punkte">Punkte</label>
        <input id="punkte" type="number" v-model.number="editData.points" v-on:keyup.enter="setFocus('maxPoints')">
      </div>
      <div class="column column-17">
        <label for="maxPoints">max.&nbsp;Punkte</label>
        <input id="maxPoints" type="number" v-model.number="editData.maxPoints" v-on:keyup.enter="enter()">
      </div>
      <div class="column column-17 full">
        <label for="pres">Vortr&auml;ge</label>
        <input id="pres" type="number" v-model.number="editData.presentations" v-on:keyup.enter="enter()">
      </div>
      <div class="column full" style="flex: 0 0 11rem">
        <label>&nbsp;</label>
        <button class="button submit-button" type="button" @click="addTask()" v-if="isNew">Hinzufügen</button>
        <button class="button submit-button" type="button" @click="saveEditTask()" v-else>Speichern</button>
      </div>
    </fieldset>
  </div>
</template>

<script>
import api from "../api";
export default {
  data() {
    return {
      gridColumns: ["Aufgabe", "Punkte", "Vorträge"],
      gridColumnsShort: ["name", "points", "maxPoints", "presentations"],
      gridData: [],
      editData: { maxPoints: 0, points: 0, name: "", presentations: 0, id: 0 },
      isNew: true,
      account: this.$route.params.account,
      course: this.$route.params.course,
      settings: false,
      percent: 70,
      perc: "-",
      pres: 0,
      showPoints: false,
      von: 100,
      zuErreichen: 70,
      vorträge: 1
    };
  },
  created() {
    api
      .GET(
        "/tasks/" +
          encodeURIComponent(this.account) +
          "/" +
          encodeURIComponent(this.course)
      )
      .then(res => {
        this.gridData = res.body || [];
      });
    api.GET("/perc/" + encodeURI(this.account) + "/" + encodeURI(this.course)).then( res => { this.perc = res.body.percentage  || "-"; });
    api.GET("/pres/" + encodeURI(this.account) + "/" + encodeURI(this.course)).then( res => { this.pres = res.body.presentations  || "-"; });
  },
  methods: {
    addTask: async function() {
      var res = await api.POST(
        "/tasks/" +
          encodeURIComponent(this.account) +
          "/" +
          encodeURIComponent(this.course),
        this.editData
      );
      this.gridData.push({ ...this.editData, id: res.body.created });
      this.editData = { maxPoints: 0, points: 0, name: "", presentations: 0 };
    },
    enter: function() {
      if (this.isNew) {
        this.addTask();
      } else {
        this.saveEditTask();
      }
    },
    setFocus: function(id) {
      document.getElementById(id).focus();
    },
    deleteItem: function(delId, id) {
      this.gridData.splice(delId, 1);
      api.DELETE(
        "/tasks/" +
          encodeURIComponent(this.account) +
          "/" +
          encodeURIComponent(this.course) +
          "/" +
          encodeURIComponent(id)
      );
    },
    edit: function(row) {
      this.isNew = false;
      for (let i in this.editData) this.editData[i] = row[i];
    },
    saveEditTask: async function() {
      //Upload changes
      var res = await api.PUT(
        "/tasks/" +
          encodeURIComponent(this.account) +
          "/" +
          encodeURIComponent(this.course) +
          "/" +
          encodeURIComponent(this.editData.id),
        this.editData
      );
      this.gridData.forEach((v, i, a) => {
        if (v.id == this.editData.id) a.splice(i, 1, this.editData);
      });

      this.isNew = true;
      this.editData = {
        maxPoints: 0,
        points: 0,
        name: "",
        presentations: 0,
        id: 0
      };
    },
    changeP: async function() {
      if (this.showPoints) {
        this.percent =
          Math.round(
            parseInt(this.von) > 0
              ? parseInt(this.zuErreichen) / parseInt(this.von) * 100
              : 0
          ) + "%";
      }
      var res = await api.PUT("/perc/" + encodeURI(this.account) + "/" + encodeURI(this.course), {'percentage': parseInt(this.percent)})
      this.perc = parseInt(this.percent);
      var res = await api.PUT("/pres/" + encodeURI(this.account) + "/" + encodeURI(this.course), {'presentations': parseInt(this.vorträge)})
      this.pres = parseInt(this.vorträge);
    }
  },
  computed: {
    gesPoints: function() {
      let res = 0;
      for (let index in this.gridData) {
        res += parseInt(this.gridData[index]["points"]);
      }
      return res;
    },
    gesMaxPoints: function() {
      let res = 0;
      for (let index in this.gridData) {
        res += parseInt(this.gridData[index]["maxPoints"]);
      }
      return res;
    },
    gesPres: function() {
      let res = 0;
      for (let index in this.gridData) {
        res += parseInt(this.gridData[index]["presentations"]);
      }
      return res;
    },
    gesPerc: function() {
      return this.gesPoints
        ? Math.round(this.gesPoints / this.gesMaxPoints * 1000) / 10 + "%"
        : 0 + "%";
    }
  }
};
</script>

<style scoped>
.right {
  position: absolute;
  right: 3rem;
  margin-top: 4px;
}
.course {
  width: 100%;
  font-weight: normal;
  font-size: 15pt;
  padding-left: 13px;
  display: flex;
  justify-content: center;
}
.grow-enter-active {
  animation: grow-up 0.9s;
}
.fade-leave-active {
  animation: grow-up 0.5s reverse;
}
@keyframes grow-up {
  0% {
    transform: scaleY(0);
  }
  75% {
    transform: scaleY(0, 85);
  }
  100% {
    transform: scaleY(1);
  }
}
.settings {
  background-color: aliceblue;
  border-radius: 7px;
  padding-bottom: 5px;
  padding-top: 5px;
  font-weight: normal;
}
h3 {
  margin-bottom: 5px;
  margin-left: 25px;
}

/*Slider*/
.switch {
  position: relative;
  display: inline-block;
  width: 30px;
  height: 14px;
  margin-top: 5px;
  margin-left: 5px;
  margin-right: 5px;
}
.switch input {
  display: none;
}
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: 0.4s;
  transition: 0.4s;
  border-radius: 34px;
}
.slider:before {
  position: absolute;
  content: "";
  height: 10px;
  width: 10px;
  left: 2px;
  bottom: 2px;
  background-color: white;
  -webkit-transition: 0.4s;
  transition: 0.4s;
  border-radius: 50%;
}
input:checked + .slider {
  background-color: #2196f3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196f3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(16px);
  -ms-transform: translateX(16px);
  transform: translateX(16px);
}
</style>

