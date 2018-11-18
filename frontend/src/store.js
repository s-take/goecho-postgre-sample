import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

const BACKEND_URL = 'http://localhost:8080';

const SET_TASKS = 'SET_TASKS';
const CREATE_TASK = 'CREATE_TASK';
const DELETE_TASK = 'DELETE_TASK';

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    tasks: [],
  },
  mutations: {
    [SET_TASKS](state, tasks) {
      state.tasks = tasks;
    },
    [CREATE_TASK](state, task) {
      state.tasks.push(task)
    },
    [DELETE_TASK](state, task) {
      state.tasks.splice(state.tasks.indexOf(task), 1)
    },
  },
  actions: {
    getTasks({ commit }) {
      axios
        .get(`${BACKEND_URL}/tasks`)
        .then(({ data }) => {
          commit(SET_TASKS, data);
        })
        //.catch((error) => console.log(error));
    },
    async createTask({ commit }, body) {
      //const body = {
      //  name: task.name
      //}
      // const { data } = await axios.post(`${BACKEND_URL}/tasks`, body)
      const { data } = await axios.post(`${BACKEND_URL}/tasks`, body)
      commit(CREATE_TASK, data);
    },
    async deleteTask({ commit }, task) {
      await axios.delete(`${BACKEND_URL}/tasks/` + task.id)
      commit(DELETE_TASK, task);
    },
  }
});

store.dispatch('getTasks');

export default store;
