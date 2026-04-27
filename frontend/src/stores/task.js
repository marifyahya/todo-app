import { defineStore } from "pinia";
import api from "../api/axios";

export const useTaskStore = defineStore("task", {
  state: () => ({
    tasks: [],
    loading: false,
    error: null,
  }),
  actions: {
    async fetchTasks(params = {}) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.get("/tasks", { params });
        this.tasks = response.data.data;
      } catch (err) {
        this.error = err.response?.data?.message || "Failed to fetch tasks";
      } finally {
        this.loading = false;
      }
    },
    async createTask(title) {
      try {
        const response = await api.post("/tasks", { title });
        this.tasks.unshift(response.data.data); // Add to top of list
        return response.data.data;
      } catch (err) {
        throw err.response?.data?.message || "Failed to create task";
      }
    },
    async updateTask(id, updates) {
      try {
        const response = await api.put(`/tasks/${id}`, updates);
        const index = this.tasks.findIndex((t) => t.id === id);
        if (index !== -1) {
          this.tasks[index] = response.data.data;
        }
      } catch (err) {
        throw err.response?.data?.message || "Failed to update task";
      }
    },
    async deleteTask(id) {
      try {
        await api.delete(`/tasks/${id}`);
        this.tasks = this.tasks.filter((t) => t.id !== id);
      } catch (err) {
        throw err.response?.data?.message || "Failed to delete task";
      }
    },
    async toggleStatus(task) {
      const newStatus = task.status === "completed" ? "pending" : "completed";
      return this.updateTask(task.id, { status: newStatus });
    },
  },
});
