<div class="spinner"></div>

<style>
  .spinner {
    width: 48px;
    height: 48px;
    border: 4px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ff3e00; /* Svelte orange */
    border-radius: 50%;
    margin-top: 7em;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
