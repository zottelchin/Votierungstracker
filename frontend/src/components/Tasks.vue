<template>
  <div id="wrapper" class="container">
    <a href="/"><h1>Votierungstracker</h1></a>
    <div class="row">
      <div class="column">
        <table>
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
                {{ entry['task'] }}
              </td>
              <td>
                {{ entry['points'] }}
              </td>
              <td>
                {{ entry['maxPoints'] }}
              </td>
              <td>
                {{ entry['pres'] }}
              </td>
              <td class="small">
                <i class="fas fa-edit" @click="edit(entry)"></i>
                <i class="fas fa-trash-alt" @click="deleteItem(index, entry['id'])"></i>
              </td>
            </tr>
            <tr>
              <td>
                <b>Gesamt</b>
              </td>
              <td>
                <b>{{ gesPoints }}</b>
              </td>
              <td>
                <b>{{ gesMaxPoints }}</b>
              </td>
              <td>
                <b>{{ gesPres }}</b>
              </td>
              <td>
                {{ gesPerc }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="column column-33">
        <fieldset>
          <label for="Aufgabe">Aufgabe</label>
          <input id="Aufgabe" type="text" v-model="editData.task">
          <label for="punkte">Punkte</label>
          <input id="punkte" type="number" v-model="editData.points">
          <label for="maxPoints">m&ouml;gliche Punkte</label>
          <input id="maxPoints" type="number" v-model="editData.maxPoints">
          <label for="pres">Vortr&auml;ge</label>
          <input id="pres" type="number" v-model="editData.pres">
          <button class="button" type="button" @click="addTask()" v-if="isNew">Hinzuf&uuml;gen</button>
          <button class="button" type="button" @click="saveEditTask()" v-else>&Auml;ndern hochladen</button>
        </fieldset>
      </div>
    </div>
  </div>
</template>

<script>
  import api from "../api"
  export default {
    data() {
      return {
          gridColumns: ['Aufgabe', 'Punkte', 'mögliche Punkte', 'Vorträge'],
          gridColumnsShort: ['task', 'points', 'maxPoints', "pres"],
          gridData: [],
          editData: { maxPoints: 0, points: 0, task: "", pres: 0, id: 0 },
          isNew: true,
          user: "",
          clas: "",
      }
    },
    created() {
      this.user = this.$route.query.user
      this.clas = this.$route.query.class
      api.get('/getVotes/' + this.user + '/' + this.clas).then(res => {
        this.gridData = res.body.items ? res.body.items : []
      })

    },
    methods: {
      addTask: async function () {
        var res = await api.get('/addVotes/'  + this.user + '/' + this.clas + '/' + this.editData.task + '/' + this.editData.points + '/' + this.editData.maxPoints + '/' + this.editData.pres)
        this.gridData.push({id: Number(res.body.created), task: this.editData.task, points: Number(this.editData.points), maxPoints: Number(this.editData.maxPoints), pres: Number(this.editData.pres)})
        this.editData = { maxPoints: 0, points: 0, task: "", pres: 0}


      },
      deleteItem: function(delId, id) {
        this.gridData.splice(delId, 1)
        api.get("/deleteVotes/" + id)
        //TODO: Im Backend löschen
      },
      edit: function(row){
        this.isNew = false
        this.editData.id = Number(row.id)
        this.editData.task = row.task
        this.editData.points = Number(row.points)
        this.editData.maxPoints = Number(row.maxPoints)
        this.editData.pres = Number(row.pres)

      },
      saveEditTask:function(){
        //Upload changes

        this.isNew = true
        this.editData = { maxPoints: 0, points: 0, task: "", pres: 0, id: 0 }
      }
    },
    computed: {
      gesPoints: function() {
        let res = 0
        for (let index in this.gridData){
          res += this.gridData[index]['points']
        }
        return res
      },
      gesMaxPoints: function() {
        let res = 0
        for (let index in this.gridData){
          res += this.gridData[index]['maxPoints']
        }
        return res
      },
        gesPres: function() {
        let res = 0
        for (let index in this.gridData){
          res += Number(this.gridData[index]['pres'])
        }
        return res
      },
      gesPerc: function() {
        return this.gesPoints ? Math.round(this.gesPoints/this.gesMaxPoints * 1000)/10 + '%' : 0 + '%'
      }
    }
  }
</script>

