<template>
  <NotificationGroup group="alerts">
    <div class="alert">
      <div class="w-full max-w-sm">
        <Notification v-slot="{ notifications }" enter="transform ease-out duration-300 transition" 
          enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0" leave="transition ease-in duration-500" 
          leave-from="opacity-100" leave-to="opacity-0" move="transition duration-500" move-delay="delay-300">
          <div v-for="notification in notifications" :key="notification.id" class="notification" :class="notificationClasses(notification.type)">
            <div class="ico" :class="iconBgColor(notification.type)">
              <SvgIcon :name="notification.type" class="text-white" />
            </div>
            <div class="message">
              <div class="mx-3">
                <span :class="textColor(notification.type)" class="font-semibold">{{ notificationType(notification.type) }}</span>
                <p>{{ notification.text }}</p>
              </div>
            </div>
          </div>
        </Notification>
      </div>
    </div>
  </NotificationGroup>
</template>

<script setup lang="ts">
import { notify, Notification, NotificationGroup } from "notiwind";
import { SvgIcon } from "@/components";
import { onBeforeUnmount } from 'vue';

const NOTIFICATION_DURATION = 4000;

const eventListeners = [];

const handleNotification = (eventName: string, notificationType: string) => {
  const callback = (e: Event) => {
    notify(
      {
        group: "alerts",
        type: notificationType,
        text: (<any>e).detail,
      },
      NOTIFICATION_DURATION
    );
  };
  
  addEventListener(eventName, callback);
  eventListeners.push({ eventName, callback });
};

onBeforeUnmount(() => {
  eventListeners.forEach(({ eventName, callback }) => {
    removeEventListener(eventName, callback);
  });
});

handleNotification("connextError", "error");
handleNotification("connextSuccess", "success");
handleNotification("connextWarning", "warning");
handleNotification("connextInfo", "info");
</script>


<style lang="scss">
.alert {
  @apply pointer-events-none fixed bottom-0 right-0 z-50 flex items-start justify-end p-6 px-4 py-6;
}

.notification {
  @apply mx-auto mt-4 flex w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-md;
}

.ico {
  @apply flex w-12 items-center justify-center;
}

.message {
  @apply -mx-3 px-4 py-2;

  p {
    @apply text-sm text-gray-600;
  }
}
</style>
