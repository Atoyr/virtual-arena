<!-- +page -->
<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  let socket: WebSocket;
  let messages: string[] = [];
  let newMessage = '';

  onMount(() => {
    // WebSocket サーバーに接続
    socket = new WebSocket('ws://localhost:8080/ws');

    socket.addEventListener('open', () => {
      console.log('WebSocket connected');
    });

    socket.addEventListener('message', (event) => {
      // サーバーからのメッセージを配列に追加
      messages = [...messages, event.data];
    });

    socket.addEventListener('close', () => {
      console.log('WebSocket disconnected');
    });

    socket.addEventListener('error', (err) => {
      console.error('WebSocket error:', err);
    });
  });

  onDestroy(() => {
    // コンポーネント破棄時に切断
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.close();
    }
  });

  function send() {
    if (socket && socket.readyState === WebSocket.OPEN && newMessage.trim() !== '') {
      socket.send(newMessage);
      newMessage = '';
    }
  }
</script>

<style>
  .chat {
    max-width: 400px;
    margin: 1rem auto;
  }
  input {
    width: 70%;
    padding: 0.5rem;
    margin-right: 0.5rem;
  }
  button {
    padding: 0.5rem 1rem;
  }
  ul {
    list-style: none;
    padding: 0;
    margin-top: 1rem;
    border: 1px solid #ccc;
    max-height: 200px;
    overflow-y: auto;
  }
  li {
    padding: 0.25rem 0.5rem;
    border-bottom: 1px solid #eee;
  }
</style>

<div class="chat">
  <h2>WebSocket チャット</h2>
  <div>
    <input
      bind:value={newMessage}
      placeholder="メッセージを入力"
      on:keydown={(e) => e.key === 'Enter' && send()}
    />
    <button on:click={send}>送信</button>
  </div>
  <ul>
    {#each messages as msg (msg + Math.random())}
      <li>{msg}</li>
    {/each}
  </ul>
</div>
