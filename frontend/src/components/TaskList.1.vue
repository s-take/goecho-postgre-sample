<template>
  <div>
    <b-field label="New Task">
      <b-input v-model="taskName"></b-input>
    </b-field>

    <b-field class="submit">
        <button class="button is-primary" @click="createTask">add</button>
    </b-field>

    <b-table
            :data="tasks"
            :loading="loading"

            paginated
            backend-pagination
            :total="total"
            :per-page="perPage"
            @page-change="onPageChange"

            backend-sorting
            :default-sort-direction="defaultSortOrder"
            :default-sort="[sortField, sortOrder]"
            @sort="onSort">

            <template slot-scope="props">
                <b-table-column field="name" label="Task Name">
                    {{ props.row.name }}
                </b-table-column>
                <b-table-column field="created_at" label="Date">
                    {{ props.row.created_at }}
                </b-table-column>
                <b-table-column field="id" label="ID" sortable>
                    {{ props.row.id | truncate(8) }}
                </b-table-column>
                <b-table-column label="Delete">
                   <button class="button is-primary" @click="deleteTask(props.row)">delete</button>️ 
                </b-table-column>
            </template>
        </b-table>

  </div>
</template>

<script>
import { mapState } from 'vuex';
export default {
  data() {
    return {
      taskName: '',
      total: 0,
      loading: false,
      sortField: 'created_at',
      sortOrder: 'desc',
      defaultSortOrder: 'desc',
      page: 1,
      perPage: 20
    };
  },
  computed: mapState({
    tasks: (state) => state.tasks,
  }),
  methods: {
    createTask() {
      if (this.taskName.length != 0) {
        this.$store.dispatch('createTask', { name: this.taskName });
        this.taskName = '';
      }
    },
    deleteTask(task) {
      this.$store.dispatch('deleteTask', task);
    },
    /*
     * Handle page-change event
     */
    onPageChange(page) {
        this.page = page
        // this.loadAsyncData()
    },
    /*
     * Handle sort event
     */
    onSort(field, order) {
        this.sortField = field
        this.sortOrder = order
        // this.loadAsyncData()
    },
  },
  filters: {
  /**
   * Filter to truncate string, accepts a length parameter
   */
    truncate(value, length) {
      return value.length > length
        ? value.substr(0, length) + '...'
        : value
    }
  },
};
</script>