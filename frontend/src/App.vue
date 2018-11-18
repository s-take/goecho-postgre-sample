<template>
  <div class="container">
      <h1>Sample frontend My Tasks</h1>
      <br>
      <TaskList />
  </div>
</template>

<script>
import TaskList from '@/components/TaskList';

export default {
  components: {
    TaskList,
  },
};
</script>

<style lang="scss" scoped>
.container {
  max-width: 768px;
}
</style>