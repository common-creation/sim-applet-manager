import { makeAutoObservable } from "mobx";
import { CheckUpdate } from '../../wailsjs/go/main/App';
import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";

export const VersionStore = () => {
  return makeAutoObservable({
    version: "",

    async checkUpdate() {
      try {
        this.version = await CheckUpdate();
        console.log(this.version);
      } catch (e) {
        console.error(e);
      }
    },

    openReleasePage() {
      BrowserOpenURL("https://github.com/common-creation/sim-applet-manager/releases/latest");
    }
  });
};