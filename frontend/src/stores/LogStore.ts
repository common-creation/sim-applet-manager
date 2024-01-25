import { makeAutoObservable } from "mobx";
import { EventsOn } from '../../wailsjs/runtime/runtime';

const logStore = () => {
  return makeAutoObservable({
    logs: [] as string[],

    async reset() {
      this.logs = [];
    },

    watch() {
      EventsOn("gpLogs", line => {
        this.logs.push(line);
        requestAnimationFrame(() => document.getElementById("eof")?.scrollIntoView());
      });
    }
  });
};

export const LogStore = logStore();
