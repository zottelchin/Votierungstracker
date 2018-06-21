<template>
  <div id="wrapper" class="container">
    <div class="back">
      <router-link :to="'/'+encodeURIComponent(account)" class="back-button"><i class="fa fa-chevron-left square-icon"></i> Zurück zu meinen Kursen</router-link>
      <div class="course">
        <span><!--aktueller Kurs:--><!--<i class="fa fa-book"></i> --><b>{{ course }}</b></span>
        <i class="settings-button fas fa-cog" @click="toggleSettings()"></i>
      </div>


      <div class="container settings" :class="{open: settingsOpen}">
        <div class="row">
          <h3><i class="fa fa-cog"></i> Einstellungen</h3>
        </div>
        <div class="row" style="margin-bottom: 0.5rem !important">
            <div style="display: flex; align-items: center;">
              <label style="font-weight: normal; margin-right: 1rem;">Benötigte&nbsp;Punkte:</label>
              <input v-if="settings.minimumType == 'points'" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" step="1" v-model="settings.minimumPoints">
              <input v-if="settings.minimumType == 'percent'" type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="100" step="0.1" v-model="settings.minimumPercent">
              <select v-model="settings.minimumType">
                <option value="none">Keine</option>
                <option value="percent">Prozent</option>
                <option value="points">Punkte</option>
              </select>
            </div>
        </div>
        <div class="row">
          <div style="display: flex; align-items: center;">
            <label style="font-weight: normal; margin-right: 1rem;">Benötigte&nbsp;Vorträge:</label>
            <input type="number" style="max-width: 65px; margin-left: 2px; display: inline;" min="0" max="10" v-model="settings.minimumPresentations">
          </div>
        </div>
        <div class="row" style="margin-top: 0.5rem !important; justify-content: flex-end">
          <button class="button" @click="saveSettings()">Speichern</button>
        </div>
      </div>


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
            <b>{{ gesPres }}</b><span v-if="minimumPresentations > 0"><span> /</span><b style="opacity:0.7; margin-left:3px;">{{ minimumPresentations }}</b></span>
          </td>
          <td>
            <span v-if="minimumType != 'none'"><strong style="margin-right:3px;">{{ gesPerc }}</strong><span>/</span><strong style="opacity:0.7; margin-left:3px;">{{ minimumType == 'percent' ? minimumPercent + "%" : minimumPoints }}</strong></span>
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
        <button class="button submit-button aniInput" type="button" @click="addTask()" v-if="isNew">Hinzufügen</button>
        <button class="button submit-button aniInput" type="button" @click="saveEditTask()" v-else>Speichern</button>
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
      minimumType: "none",
      minimumPoints: 50,
      minimumPercent: 50,
      minimumPresentations: 1,
      settingsOpen: false,
      settings: {
        minimumType: "none",
        minimumPoints: 50,
        minimumPercent: 50,
        minimumPresentations: 1
      }
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
    api.GET("/perc/" + encodeURI(this.account) + "/" + encodeURI(this.course)).then( res => { this.minimumType = res.body.type || "none"; this.minimumPoints = this.minimumType == "points" ? res.body.minimum : 50; this.minimumPercent = this.minimumType == "percent" ? res.body.minimum : 50; });
    api.GET("/pres/" + encodeURI(this.account) + "/" + encodeURI(this.course)).then( res => { this.minimumPresentations = res.body.presentations || 0; });
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
    toggleSettings: function() {
      this.settingsOpen = !this.settingsOpen;
      if (this.settingsOpen) {
        for (var i in this.settings) {
          console.log("Setting:",i);
          this.$set(this.settings, i, this[i]);
        }
      }
    },
    saveSettings: async function() {
      var res = await api.PUT("/perc/" + encodeURI(this.account) + "/" + encodeURI(this.course), { type: this.settings.minimumType, minimum: this.settings.minimumType == "percent" ? parseFloat(this.settings.minimumPercent) :  parseInt(this.settings.minimumPoints) })
      this.minimumType = this.settings.minimumType;
      this.minimumPercent = parseFloat(this.settings.minimumPercent);
      this.minimumPoints = parseInt(this.settings.minimumPoints);
      var res = await api.PUT("/pres/" + encodeURI(this.account) + "/" + encodeURI(this.course), { presentations: parseInt(this.settings.minimumPresentations) })
      this.minimumPresentations = parseInt(this.settings.minimumPresentations);
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
      if (this.minimumType == "points") return this.gesPoints;
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
  margin-top: 0.5rem;
  font-weight: normal;
  font-size: 2.5rem;
  display: flex;
}
.course>span>i { margin-right: 0.5rem; }
.settings-button { margin-left: 1rem; margin-top: 1.15rem; font-size: 1.5rem; color: #606c76; cursor: pointer; }
.back-button { color: #606c76; }
.settings-button:hover, .back-button:hover { color: #9b4dca; }
.settings {
  background-color: #eeeeee; border: 2px solid #bbb;
  margin: 0 -1rem -30px; width: auto; overflow: hidden;
  border-radius: 3px;
  padding: 15px !important; box-sizing: border-box;
  font-weight: normal;
  transition: max-height 0.4s, opacity 0.4s, margin-top 0.4s, margin-bottom 0.4s;
  max-height: 0; opacity: 0;
}
.settings>.row { margin: 0 !important; width: 100%; }
.open {
  opacity: 1;
  margin-top: 10px; margin-bottom: 0;
  max-height: 350px;
}
.settings>h3 {
  margin-bottom: 15px;
}
.settings input, .settings select { background-color: #fff !important; margin-bottom: 0; width: auto !important; }
.settings label { margin-bottom: 0; display: inline-block !important; }
.settings button { margin-bottom: 0; }

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

