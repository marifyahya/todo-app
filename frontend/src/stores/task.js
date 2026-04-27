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
        return response.data;
      } catch (err) {
        this.error = err.response?.data?.message || "Failed to fetch tasks";
        throw err;
      } finally {
        this.loading = false;
      }
    },
    async createTask(title) {
      try {
        const response = await api.post("/tasks", { title });
        this.tasks.unshift(response.data.data);
        return response.data;
      } catch (err) {
        throw err;
      }
    },
    async updateTask(id, updates) {
      try {
        const response = await api.put(`/tasks/${id}`, updates);
        const index = this.tasks.findIndex((t) => t.id === id);
        if (index !== -1) {
          this.tasks[index] = response.data.data;
        }
        return response.data;
      } catch (err) {
        throw err;
      }
    },
    async deleteTask(id) {
      try {
        const response = await api.delete(`/tasks/${id}`);
        this.tasks = this.tasks.filter((t) => t.id !== id);
        return response.data;
      } catch (err) {
        throw err;
      }
    },
    async toggleStatus(task) {
      const newStatus = task.status === "completed" ? "pending" : "completed";
      return this.updateTask(task.id, { status: newStatus });
    },
  },
});
