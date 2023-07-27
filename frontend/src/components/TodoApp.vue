<template>
  <div class="caintaner">
    <!-- heading -->
    <h1 class="text-center"> Todo App</h1>
    <!-- input -->
    <div>
      <input type="text" v-model="task" placeholder="Enter here.."/>
      <button class="add-btn" @click="SubmitTask">ADD</button>
    </div>
    <!-- table -->
    <table border={3}>
      <thead>
      <tr>
        <th>ID</th>
        <th>Task</th>
        <th>Delete</th>
        <th>Edit</th>

      </tr>
      </thead>
      <tbody>
      <tr v-for="(task,index) in tasks" :key="index">
        <td>{{ task.ID }}</td>
        <td>{{ task.task }}</td>

        <td>
          <button class="del-btn" @click="deleteTask(index)">Delete</button>
        </td>
        <td>
          <button class="edit-btn" @click="EditTask(index)">Edit</button>
        </td>
      </tr>

      </tbody>
    </table>

  </div>
</template>

<script>
import axios from 'axios'
const API_URL = "http://localhost:3000"

export default {
  name: "TodoApp",
  props: {
    msg: String,
  },
  data() {
    return {
      task: "",
      editTask: null,
      tasks: []
    }
  },
  mounted() {
    this.getTasks()
  },
  methods: {
    getTasks() {
      axios
          .get(`${API_URL}/todos`)
          .then(response => (this.tasks = response.data))
          .catch(error => console.log(error));
    },

    addTask(taskName) {
      axios
          .post(`${API_URL}/todos`, {
            task: taskName
          })
          .then(response => (this.tasks.push(response.data)))
          .catch(error => console.log(error));
    },

    // delete task
    deleteTask(index) {
      let itemId = this.tasks[index].ID

      axios
          .delete(`${API_URL}/todos/${itemId}`)
          .then(response => (this.tasks.push(response.data)))
          .catch(error => console.log(error));

      this.tasks.splice(index, 1)
    },

    // edit task
    EditTask(index) {
      this.task = this.tasks[index].task
      this.editTask = index
    },

    // Add Task
    SubmitTask() {
      if (this.task.length === 0) {
        return;
      }
      if (this.editTask != null) {
        this.tasks[this.editTask].task = this.task;

        let itemId = this.tasks[this.editTask].ID
        axios
            .put(`${API_URL}/todos/${itemId}`, {
              task: this.task
            })
            .then(response => (this.tasks.push(response.data)))
            .catch(error => console.log(error));

        this.editTask = null

      } else {
        this.addTask(this.task)
      }
    }

  }

}
</script>

<style scoped>
.caintaner {
  width: 600px;
  margin: auto;

}

table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
  margin-top: 20px;
}

td,
th {
  border: 1px solid #dddddd;
  text-align: center;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}

.add-btn {
  border: none;
  width: 100px;
  height: 30px;
  padding: 2px;
  background-color: teal;
  color: white;
}

.del-btn {
  border: none;
  background-color: red;
  color: white;
}

.edit-btn {
  border: none;
  background-color: #77b631;
  color: white;

}
</style>
