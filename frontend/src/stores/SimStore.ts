import { makeAutoObservable } from "mobx";
import { GetGpPath, ListCardReader, FetchSimInfo, SaveSimConfig, ShowErrorDialog, ListApplets, UninstallApplet, ShowConfirmDialog, InstallApplet } from '../../wailsjs/go/main/App';
import { db, gp, main } from '../../wailsjs/go/models';

export const SimStore = () => {
  return makeAutoObservable({
    gpPath: "",
    cardReaders: [] as string[],
    iccid: "",
    keys: [] as db.Key[],
    editKeys: [] as db.Key[],

    applets: [] as gp.ListResult[],

    selectedCardReaderIndex: 0,
    openSimConfig: false,

    async fetchGpPath() {
      try {
        document.body.classList.add("loading");
        this.gpPath = await GetGpPath();
      } catch (e) {
        console.error(e);
      } finally {
        document.body.classList.remove("loading");
      }
    },

    async listCardReader() {
      try {
        document.body.classList.add("loading");
        this.cardReaders = await ListCardReader();
        this.selectedCardReaderIndex = 0;

        await this.fetchSimInfo(this.cardReaders[0]);
      } catch (e) {
        console.error(e);
      } finally {
        document.body.classList.remove("loading");
      }
    },

    async setSelected(selectedIndex: number) {
      this.selectedCardReaderIndex = selectedIndex;
      await this.fetchSimInfo(this.cardReaders[selectedIndex]);
    },

    async fetchSimInfo(cardReader: string): Promise<main.SimInfo> {
      try {
        document.body.classList.add("loading");
        const { iccid, config } = await FetchSimInfo(cardReader);
        console.log(iccid, config);
        this.iccid = iccid;
        this.keys = config?.keys || [];

        if (this.keys.length > 0) {
          await this.listApplets(this.keys[0]);
        } else {
          this.applets = [];
        }
      } catch (e) {
        console.error(e);
      } finally {
        document.body.classList.remove("loading");
      }
      return {} as main.SimInfo;
    },

    async setOpenSimConfig(open: boolean) {
      if (open) {
        this.editKeys = JSON.parse(JSON.stringify(this.keys));
      } else {
        this.editKeys = [];
      }
      this.openSimConfig = open;
    },

    async saveSimConfig() {
      try {
        document.body.classList.add("loading");

        const ok = await SaveSimConfig(this.iccid, new db.Sim({ keys: this.editKeys }));
        if (!ok) {
          throw new Error("設定の保存に失敗しました");
        }
        this.keys = this.editKeys;
        this.openSimConfig = false;
      } catch (e) {
        console.log(e);
        await ShowErrorDialog("エラー", "設定の保存に失敗しました");
      } finally {
        document.body.classList.remove("loading");
      }
    },

    async listApplets(key: db.Key) {
      try {
        document.body.classList.add("loading");

        this.applets = await ListApplets(this.cardReaders[this.selectedCardReaderIndex], new db.Key(key));
      } catch (e) {
        console.log(e);
        await ShowErrorDialog("エラー", "アプレットの取得に失敗しました");
      } finally {
        document.body.classList.remove("loading");
      }
    },

    async uninstallApplet(key: db.Key, aid: string) {
      try {
        if (await ShowConfirmDialog("確認", `${aid} を削除しますか?`) !== "OK") {
          return;
        }

        document.body.classList.add("loading");

        const result = await UninstallApplet(this.cardReaders[this.selectedCardReaderIndex], new db.Key(key), aid);
        if (!result.success) {
          throw new Error("アプレットの削除に失敗しました");
        }
        await this.listApplets(key);
      } catch (e) {
        console.log(e);
        await ShowErrorDialog("エラー", "アプレットの削除に失敗しました");
      } finally {
        document.body.classList.remove("loading");
      }
    },

    async installApplet(key: db.Key, path: string, params: string): Promise<boolean> {
      try {
        document.body.classList.add("loading");

        const result = await InstallApplet(this.cardReaders[this.selectedCardReaderIndex], new db.Key(key), path, params);
        if (!result.success) {
          throw new Error("アプレットのインストールに失敗しました");
        }
        await this.listApplets(key);
        return true;
      } catch (e) {
        console.log(e);
        await ShowErrorDialog("エラー", "アプレットのインストールに失敗しました");
      } finally {
        document.body.classList.remove("loading");
      }
      return false;
    },
  });
};