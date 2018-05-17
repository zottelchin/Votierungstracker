<template>
  <div id="wrapper" class="container">
    <div class="back">
      <div class="course">
        <span><!--aktueller Kurs:--> <b>{{ course }}</b></span>
      </div>
      <router-link :to="'/'+encodeURIComponent(account)"><i class="fa fa-chevron-left"></i> Zurück zu meinen Kursen</router-link>
      <i class="fas fa-cog right"></i>
      
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
            <b>{{ gesPres }}</b>
          </td>
          <td>
            {{ gesPerc }}
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
  import api from "../api"
  export default {
    data() {
      return {
          gridColumns: ['Aufgabe', 'Punkte', 'Vorträge'],
          gridColumnsShort: ['name', 'points', 'maxPoints', "presentations"],
          gridData: [],
          editData: { maxPoints: 0, points: 0, name: "", presentations: 0, id: 0 },
          isNew: true,
          account: this.$route.params.account,
          course: this.$route.params.course,
      }
    },
    created() {
      api.GET('/tasks/' + encodeURIComponent(this.account) + '/' + encodeURIComponent(this.course)).then(res => {
        this.gridData = res.body || []
      })
    },
    methods: {
      addTask: async function () {
        var res = await api.POST('/tasks/' + encodeURIComponent(this.account) + '/' + encodeURIComponent(this.course), this.editData)
        this.gridData.push({ ...this.editData, id: res.body.created })
        this.editData = { maxPoints: 0, points: 0, name: "", presentations: 0 }
      },
      enter: function(){
        if (this.isNew){
          this.addTask()
        }else{
          this.saveEditTask()
        }
      },
      setFocus: function(id){
        document.getElementById(id).focus()
      },
      deleteItem: function(delId, id) {
        this.gridData.splice(delId, 1)
        api.DELETE("/tasks/" + encodeURIComponent(this.account) + '/' + encodeURIComponent(this.course) + '/' + encodeURIComponent(id))
      },
      edit: function(row){
        this.isNew = false
        for (let i in this.editData) this.editData[i] = row[i];
      },
      saveEditTask: async function() {
        //Upload changes
        var res = await api.PUT('/tasks/' + encodeURIComponent(this.account) + '/' + encodeURIComponent(this.course) + '/' + encodeURIComponent(this.editData.id), this.editData)
        this.gridData.forEach((v, i, a) => { if (v.id == this.editData.id) a.splice(i, 1, this.editData); })

        this.isNew = true
        this.editData = { maxPoints: 0, points: 0, name: "", presentations: 0, id: 0 }
      }
    },
    computed: {
      gesPoints: function() {
        let res = 0
        for (let index in this.gridData){
          res += parseInt(this.gridData[index]['points'])
        }
        return res
      },
      gesMaxPoints: function() {
        let res = 0
        for (let index in this.gridData){
          res += parseInt(this.gridData[index]['maxPoints'])
        }
        return res
      },
      gesPres: function() {
        let res = 0
        for (let index in this.gridData){
          res += parseInt(this.gridData[index]['presentations'])
        }
        return res
      },
      gesPerc: function() {
        return this.gesPoints ? Math.round(this.gesPoints/this.gesMaxPoints * 1000)/10 + '%' : 0 + '%'
      }
    }
  }
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
</style>

