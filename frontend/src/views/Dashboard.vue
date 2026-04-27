<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";
import { useTaskStore } from "../stores/task";

const router = useRouter();
const authStore = useAuthStore();
const taskStore = useTaskStore();

const newTaskTitle = ref("");
const selectedTask = ref(null);
const editForm = ref({
  title: "",
  description: "",
  priority: "",
  due_date: "",
});

const showDeleteModal = ref(false);
const taskToDelete = ref(null);
const isSidebarOpen = ref(false);
const activeFilter = ref({ type: "status", value: "pending" });

// Toast Notification State
const toast = ref({
  show: false,
  message: "",
  type: "success",
});

const showToast = (message, type = "success") => {
  toast.value = { show: true, message, type };
  setTimeout(() => {
    toast.value.show = false;
  }, 3000);
};

onMounted(() => {
  taskStore.fetchTasks({ status: "pending" });
});

// Watch for selectedTask changes to populate editForm
watch(selectedTask, (newVal) => {
  if (newVal) {
    editForm.value = {
      title: newVal.title,
      description: newVal.description || "",
      priority: newVal.priority || "",
      due_date: newVal.due_date
        ? new Date(newVal.due_date).toISOString().split("T")[0]
        : "",
    };
  }
});

const handleLogout = () => {
  authStore.logout();
  router.push("/login");
};

const handleAddTask = async () => {
  if (!newTaskTitle.value.trim()) return;
  try {
    const res = await taskStore.createTask(newTaskTitle.value);
    newTaskTitle.value = "";
    showToast(res.message || "Task created successfully!");
  } catch (err) {
    showToast(err.response?.data?.message || "Failed to create task", "error");
  }
};

const handleUpdateTaskDetail = async () => {
  if (!selectedTask.value) return;
  try {
    const res = await taskStore.updateTask(selectedTask.value.id, {
      ...editForm.value,
      priority: editForm.value.priority || null,
      due_date: editForm.value.due_date
        ? new Date(editForm.value.due_date)
        : null,
    });
    showToast(res.message || "Task updated successfully!");
  } catch (err) {
    showToast(err.response?.data?.message || "Failed to update task", "error");
  }
};

const handleToggle = async (task) => {
  try {
    const res = await taskStore.toggleStatus(task);
    showToast(res.message || "Status updated!");
  } catch (err) {
    showToast("Failed to update status", "error");
  }
};

const setFilter = (type, value) => {
  activeFilter.value = { type, value };
  const params = {};
  if (type === "all") {
    // If 'all' is clicked, we actually want only pending as per user request
    activeFilter.value = { type: "status", value: "pending" };
    params.status = "pending";
  } else if (type !== "all") {
    params[type] = value;
  }
  taskStore.fetchTasks(params);
  isSidebarOpen.value = false;
};

const confirmDelete = (id) => {
  taskToDelete.value = id;
  showDeleteModal.value = true;
};

const handleDelete = async () => {
  if (taskToDelete.value) {
    try {
      const res = await taskStore.deleteTask(taskToDelete.value);
      if (selectedTask.value?.id === taskToDelete.value)
        selectedTask.value = null;
      showDeleteModal.value = false;
      taskToDelete.value = null;
      showToast(res.message || "Task deleted");
    } catch (err) {
      showToast("Failed to delete task", "error");
    }
  }
};

const getPriorityClass = (priority) => {
  if (!priority) return "bg-gray-50 text-gray-400 border border-gray-100";
  switch (priority) {
    case "high":
      return "bg-red-50 text-red-600 border border-red-100";
    case "medium":
      return "bg-blue-50 text-blue-600 border border-blue-100";
    case "low":
      return "bg-gray-50 text-gray-600 border border-gray-100";
    default:
      return "bg-gray-50 text-gray-600";
  }
};

const userName = computed(() => authStore.user?.name || "User");
const userEmail = computed(() => authStore.user?.email || "user@example.com");
</script>

<template>
  <div
    class="flex h-screen bg-white font-sans text-gray-800 overflow-hidden relative"
  >
    <transition name="toast">
      <div
        v-if="toast.show"
        :class="[
          toast.type === 'success'
            ? 'bg-emerald-50 border-emerald-100 text-emerald-900 shadow-emerald-100'
            : 'bg-red-50 border-red-100 text-red-900 shadow-red-100',
        ]"
        class="fixed top-8 left-1/2 -translate-x-1/2 z-[200] flex items-center gap-3 px-6 py-3.5 rounded-2xl shadow-2xl border min-w-[300px] transition-all duration-300"
      >
        <div
          :class="toast.type === 'success' ? 'bg-emerald-500' : 'bg-red-500'"
          class="w-5 h-5 rounded-full flex items-center justify-center text-white"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-3 w-3"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              v-if="toast.type === 'success'"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="4"
              d="M5 13l4 4L19 7"
            />
            <path
              v-else
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="4"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </div>
        <span class="text-xs font-black uppercase tracking-widest">{{
          toast.message
        }}</span>
      </div>
    </transition>

    <!-- MOBILE OVERLAY -->
    <div
      v-if="isSidebarOpen"
      @click="isSidebarOpen = false"
      class="fixed inset-0 bg-black/20 backdrop-blur-sm z-40 lg:hidden"
    ></div>

    <!-- LEFT SIDEBAR -->
    <aside
      class="fixed inset-y-0 left-0 w-64 bg-gray-50 border-r border-gray-200 flex flex-col z-50 transform transition-transform duration-300 lg:relative lg:translate-x-0 flex-shrink-0"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="p-6">
        <div class="flex items-center gap-3 mb-8">
          <div
            class="w-8 h-8 bg-black rounded-lg flex items-center justify-center shadow-md"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 text-white"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 002 2h2a2 2 0 002-2"
              />
            </svg>
          </div>
          <h1 class="text-sm font-black tracking-tighter text-black uppercase">
            TaskFlow
          </h1>
        </div>

        <nav class="space-y-6">
          <div>
            <h3
              class="text-[10px] font-black text-gray-400 uppercase tracking-widest mb-3 ml-2"
            >
              Navigation
            </h3>
            <div class="space-y-1">
              <div
                @click="setFilter('status', 'pending')"
                :class="
                  activeFilter.value === 'pending'
                    ? 'bg-white border-gray-200 text-black shadow-sm'
                    : 'text-gray-400 hover:text-black'
                "
                class="flex items-center justify-between p-2.5 rounded-xl border border-transparent cursor-pointer transition-all text-sm font-bold"
              >
                <div class="flex items-center gap-3">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M4 6h16M4 12h16M4 18h7"
                    />
                  </svg>
                  <span>All Tasks</span>
                </div>
              </div>
              <div
                @click="setFilter('status', 'completed')"
                :class="
                  activeFilter.value === 'completed'
                    ? 'bg-white border-gray-200 text-black shadow-sm'
                    : 'text-gray-400 hover:text-black'
                "
                class="flex items-center gap-3 p-2.5 rounded-xl border border-transparent cursor-pointer transition-all text-sm font-bold"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"
                  />
                </svg>
                <span>Completed</span>
              </div>
            </div>
          </div>

          <div>
            <h3
              class="text-[10px] font-black text-gray-400 uppercase tracking-widest mb-3 ml-2"
            >
              Priority
            </h3>
            <div class="space-y-1">
              <div
                @click="setFilter('priority', 'high')"
                :class="
                  activeFilter.value === 'high'
                    ? 'bg-white border-gray-200 text-black shadow-sm'
                    : 'text-gray-400 hover:text-black'
                "
                class="flex items-center gap-3 p-2.5 rounded-xl border border-transparent cursor-pointer transition-all text-sm font-bold"
              >
                <div class="w-2.5 h-2.5 rounded-full bg-red-500"></div>
                <span>High Priority</span>
              </div>
              <div
                @click="setFilter('priority', 'medium')"
                :class="
                  activeFilter.value === 'medium'
                    ? 'bg-white border-gray-200 text-black shadow-sm'
                    : 'text-gray-400 hover:text-black'
                "
                class="flex items-center gap-3 p-2.5 rounded-xl border border-transparent cursor-pointer transition-all text-sm font-bold"
              >
                <div class="w-2.5 h-2.5 rounded-full bg-blue-500"></div>
                <span>Medium Priority</span>
              </div>
              <div
                @click="setFilter('priority', 'low')"
                :class="
                  activeFilter.value === 'low'
                    ? 'bg-white border-gray-200 text-black shadow-sm'
                    : 'text-gray-400 hover:text-black'
                "
                class="flex items-center gap-3 p-2.5 rounded-xl border border-transparent cursor-pointer transition-all text-sm font-bold"
              >
                <div class="w-2.5 h-2.5 rounded-full bg-gray-300"></div>
                <span>Low Priority</span>
              </div>
            </div>
          </div>
        </nav>
      </div>

      <div class="mt-auto p-4 bg-white border-t border-gray-200">
        <div
          class="flex items-center gap-3 mb-4 p-2 bg-gray-50 rounded-2xl border border-gray-100"
        >
          <div
            class="w-10 h-10 rounded-xl bg-black flex items-center justify-center text-white font-black text-sm shadow-sm"
          >
            {{ userName.charAt(0).toUpperCase() }}
          </div>
          <div class="flex-grow min-w-0">
            <p class="text-xs font-black text-black truncate">{{ userName }}</p>
            <p class="text-[10px] text-gray-400 truncate">{{ userEmail }}</p>
          </div>
        </div>
        <button
          @click="handleLogout"
          class="w-full flex items-center justify-center gap-2 py-2.5 px-4 bg-gray-50 hover:bg-red-50 hover:text-red-600 rounded-xl text-[10px] font-black text-gray-500 transition-all uppercase tracking-widest border border-gray-100"
        >
          Logout
        </button>
      </div>
    </aside>

    <!-- MAIN CONTENT WRAPPER -->
    <div class="flex flex-grow overflow-hidden">
      <!-- TASK LIST -->
      <main
        @click="selectedTask = null"
        class="flex-grow overflow-y-auto bg-white transition-all duration-300 relative"
      >
        <!-- CLICKABLE OVERLAY WHEN DETAIL OPEN (Mobile/Desktop) -->
        <div
          v-if="selectedTask"
          class="absolute inset-0 z-10 lg:hidden"
          @click.stop="selectedTask = null"
        ></div>

        <div class="max-w-3xl mx-auto px-6 py-8 relative z-0" @click.stop>
          <div class="flex items-center justify-between mb-8 lg:hidden">
            <button
              @click="isSidebarOpen = true"
              class="p-2 -ml-2 text-gray-400 hover:text-black"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
            <h1 class="text-sm font-black tracking-tighter uppercase">
              TaskFlow
            </h1>
            <div class="w-10"></div>
          </div>

          <header
            class="mb-8 border-b border-gray-100 pb-6 flex items-baseline justify-between"
          >
            <div>
              <h2
                class="text-3xl font-black text-black tracking-tighter capitalize"
              >
                {{ activeFilter.value || "All Tasks" }}
              </h2>
              <p class="text-xs text-gray-400 font-medium mt-1">
                Found {{ taskStore.tasks.length }} items
              </p>
            </div>
          </header>

          <!-- Simplified Add Task Input -->
          <div
            class="bg-gray-50 p-1.5 rounded-2xl border border-gray-200 mb-8 flex items-center gap-2 focus-within:border-black focus-within:bg-white transition-all shadow-sm"
          >
            <div class="flex-grow flex items-center gap-2 pl-3">
              <input
                v-model="newTaskTitle"
                @keyup.enter="handleAddTask"
                type="text"
                placeholder="Add a new task title..."
                class="w-full py-3 text-sm bg-transparent focus:outline-none placeholder-gray-300 font-medium"
              />
            </div>
            <button
              @click="handleAddTask"
              class="p-2.5 bg-black rounded-xl text-white hover:bg-gray-800 transition-all shadow-lg"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="3"
                  d="M12 4v16m8-8H4"
                />
              </svg>
            </button>
          </div>

          <!-- Task List -->
          <div class="space-y-2">
            <div v-if="taskStore.loading" class="flex justify-center py-12">
              <div
                class="w-6 h-6 border-3 border-gray-100 border-t-black rounded-full animate-spin"
              ></div>
            </div>

            <div
              v-else
              v-for="task in taskStore.tasks"
              :key="task.id"
              @click="selectedTask = task"
              :class="
                selectedTask?.id === task.id
                  ? 'border-black bg-gray-50'
                  : 'border-gray-100 bg-white hover:border-gray-300'
              "
              class="group p-4 rounded-xl border transition-all flex items-center justify-between cursor-pointer"
            >
              <div class="flex items-center gap-4 flex-grow min-w-0">
                <div
                  @click.stop="handleToggle(task)"
                  class="flex-shrink-0 w-5 h-5 rounded-md border-2 flex items-center justify-center transition-all"
                  :class="
                    task.status === 'completed'
                      ? 'bg-black border-black'
                      : 'border-gray-200 group-hover:border-black'
                  "
                >
                  <svg
                    v-if="task.status === 'completed'"
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-3 w-3 text-white"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="3"
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                </div>

                <div class="flex flex-col min-w-0 justify-center">
                  <span
                    :class="{
                      'line-through text-gray-300 font-medium':
                        task.status === 'completed',
                      'text-black font-bold': task.status !== 'completed',
                    }"
                    class="text-sm lg:text-base transition-all truncate leading-none"
                    >{{ task.title }}</span
                  >
                  <div
                    v-if="task.priority || task.due_date"
                    class="flex items-center gap-2 mt-1"
                  >
                    <span
                      v-if="task.priority"
                      :class="getPriorityClass(task.priority)"
                      class="text-[8px] font-black uppercase tracking-widest px-1.5 py-0.5 rounded-md"
                      >{{ task.priority }}</span
                    >
                    <span
                      v-if="task.due_date"
                      class="text-[10px] text-gray-400 font-medium"
                    >
                      {{ new Date(task.due_date).toLocaleDateString() }}
                    </span>
                  </div>
                </div>
              </div>

              <button
                @click.stop="confirmDelete(task.id)"
                class="p-1.5 text-gray-200 hover:text-black opacity-0 group-hover:opacity-100 transition-all"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                  />
                </svg>
              </button>
            </div>
          </div>

          <div
            v-if="!taskStore.loading && taskStore.tasks.length === 0"
            class="text-center py-20 border-2 border-dashed border-gray-50 rounded-2xl"
          >
            <p
              class="text-gray-300 font-black text-sm uppercase tracking-tighter"
            >
              No tasks found.
            </p>
          </div>
        </div>
      </main>

      <!-- RIGHT DETAIL PANE -->
      <aside
        class="fixed inset-y-0 right-0 w-full sm:w-80 bg-gray-50 border-l border-gray-200 flex flex-col z-[55] transform transition-transform duration-300 lg:relative lg:translate-x-0 lg:fixed-none flex-shrink-0 shadow-2xl lg:shadow-none"
        :class="selectedTask ? 'translate-x-0' : 'translate-x-full lg:hidden'"
      >
        <div
          class="p-5 border-b border-gray-200 flex items-center justify-between bg-white"
        >
          <h3
            class="text-[10px] font-black uppercase tracking-widest text-black"
          >
            Task Details
          </h3>
          <button
            @click="selectedTask = null"
            class="text-gray-400 hover:text-black p-1"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        <div class="flex-grow overflow-y-auto p-5 space-y-6">
          <div>
            <label
              class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"
              >Title</label
            >
            <input
              v-model="editForm.title"
              type="text"
              class="w-full bg-transparent border-none p-0 text-lg font-black text-black focus:ring-0 placeholder-gray-200"
            />
          </div>

          <div>
            <label
              class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"
              >Description</label
            >
            <textarea
              v-model="editForm.description"
              rows="5"
              placeholder="Add notes here..."
              class="w-full bg-white border border-gray-200 rounded-xl p-3 text-xs font-medium focus:outline-none focus:border-black transition-all resize-none shadow-sm"
            ></textarea>
          </div>

          <div class="space-y-5">
            <div>
              <label
                class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-3 ml-1"
                >Priority</label
              >
              <div class="grid grid-cols-4 gap-2">
                <button
                  v-for="p in ['', 'low', 'medium', 'high']"
                  :key="p"
                  @click="editForm.priority = p"
                  :class="[
                    editForm.priority === p
                      ? p === 'high'
                        ? 'bg-red-500 text-white border-red-500 shadow-lg shadow-red-100'
                        : p === 'medium'
                          ? 'bg-blue-500 text-white border-blue-500 shadow-lg shadow-blue-100'
                          : p === 'low'
                            ? 'bg-gray-800 text-white border-gray-800 shadow-lg shadow-gray-100'
                            : 'bg-black text-white border-black shadow-lg shadow-gray-100'
                      : 'bg-white text-gray-400 border-gray-100 hover:border-gray-300',
                  ]"
                  class="py-2.5 rounded-xl border text-[10px] font-black uppercase tracking-widest transition-all"
                >
                  {{ p || "None" }}
                </button>
              </div>
            </div>
            <div>
              <label
                class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-3 ml-1"
                >Due Date</label
              >
              <div class="relative group">
                <div
                  class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 group-focus-within:text-black transition-colors"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                    />
                  </svg>
                </div>
                <input
                  v-model="editForm.due_date"
                  type="date"
                  class="w-full bg-white border border-gray-100 rounded-2xl pl-11 pr-4 py-3.5 text-xs font-bold text-black focus:border-black focus:ring-0 transition-all shadow-sm cursor-pointer"
                />
              </div>
            </div>
          </div>
        </div>

        <div class="p-5 border-t border-gray-200 bg-white">
          <button
            @click="handleUpdateTaskDetail"
            class="w-full py-3 bg-black text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-gray-200 hover:bg-gray-800 transition-all"
          >
            Save Changes
          </button>
        </div>
      </aside>
    </div>

    <!-- CUSTOM DELETE MODAL -->
    <div
      v-if="showDeleteModal"
      @click="showDeleteModal = false"
      class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm animate-in fade-in duration-200"
    >
      <div
        @click.stop
        class="bg-white rounded-3xl shadow-2xl max-w-[320px] w-full p-8 text-center animate-in zoom-in-95 duration-200"
      >
        <div
          class="w-16 h-16 bg-gray-50 rounded-2xl flex items-center justify-center text-black mx-auto mb-6 border border-gray-100"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            />
          </svg>
        </div>
        <h3 class="text-xl font-black text-black mb-1">Delete Task?</h3>
        <p class="text-xs text-gray-500 mb-8 font-medium">
          This cannot be undone.
        </p>
        <div class="flex gap-2">
          <button
            @click="showDeleteModal = false"
            class="flex-grow py-3 bg-gray-100 rounded-xl font-bold text-[10px] text-gray-600 hover:bg-gray-200 transition-all uppercase tracking-widest"
          >
            Cancel
          </button>
          <button
            @click="handleDelete"
            class="flex-grow py-3 bg-black rounded-xl font-bold text-[10px] text-white hover:bg-gray-800 transition-all uppercase tracking-widest"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800;900&display=swap");

body {
  font-family: "Inter", sans-serif;
  letter-spacing: -0.01em;
}

::-webkit-scrollbar {
  width: 4px;
}
::-webkit-scrollbar-track {
  background: transparent;
}
::-webkit-scrollbar-thumb {
  background: #f1f1f1;
  border-radius: 10px;
}

/* Toast Animation */
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}
.toast-enter-from {
  opacity: 0;
  transform: translate(-50%, -20px);
}
.toast-leave-to {
  opacity: 0;
  transform: translate(-50%, -20px);
}

.animate-in {
  animation-duration: 200ms;
  animation-fill-mode: both;
}
.fade-in {
  animation-name: fade-in;
}
.zoom-in-95 {
  animation-name: zoom-in-95;
}

@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
@keyframes zoom-in-95 {
  from {
    transform: scale(0.95);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}
</style>
