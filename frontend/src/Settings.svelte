<script lang="ts">
  import { onMount, createEventDispatcher } from "svelte";
  import { Get, Save } from "../wailsjs/go/userconfig/ConfigStore.js";
  import { userconfig } from "../wailsjs/go/models.js";

  const dispatch = createEventDispatcher();

  function closeSettings() {
    dispatch("close");
  }

  let config: userconfig.Config;
  let RAKey: string;
  let RAText: string;

  onMount(async () => {
    window.addEventListener("keydown", (e) => {
      if (e.key == "Escape") {
        closeSettings();
      }
    });

    let userconfig = await Get();
    console.log(userconfig);
    RAKey = userconfig.RetroachievmentKey;
  });

  async function handleSubmit() {
    RAKey = RAText;
    await Save({ ...config, RetroachievmentKey: RAText });
  }

  async function discardRAKey() {
    RAKey = "";
    Save({ ...config, RetroachievmentKey: "" });
  }
</script>

<!-- TODO: make this beautiful -->
<dialog class="nosetelect">
  <h1>Settings</h1>
  <div class="settings-content">
    {#if RAKey}
      <div>
        <p>Retroachievemnts Api key: {RAKey}</p>
        <button on:click|preventDefault={discardRAKey}>Discard</button>
      </div>
    {:else}
      <p>Please Provide Api key:</p>
      <form on:submit|preventDefault={handleSubmit}>
        <input
          type="text"
          bind:value={RAText}
          placeholder="Type something..."
        />
        <button type="submit" on:submit={handleSubmit}>Submit</button>
      </form>
    {/if}
  </div>
</dialog>
