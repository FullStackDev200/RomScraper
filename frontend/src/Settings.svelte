<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { NewConfigStore, GetConfig } from "../wailsjs/go/main/App.js";

  const dispatch = createEventDispatcher();

  function closeSettings() {
    dispatch("close");
  }

  onMount(async () => {
    window.addEventListener("keydown", (e) => {
      if (e.key == "Escape") {
        closeSettings();
      }
    });

    let cfs = await NewConfigStore();
    console.log(cfs);
    let cfg = GetConfig(cfs);
    console.log(cfg);
  });
</script>

<dialog class="nosetelect">
  <h1>Settings</h1>
  <div class="settings-content">
    <p>Retroachievemnts Api key:</p>
  </div>
</dialog>
