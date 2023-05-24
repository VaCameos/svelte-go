<script>
    import { goto } from '$app/navigation';
    import { createUserAjax, loginAjax } from '../../util/api';
    import Snackbar, {  Label } from '@smui/snackbar';
    // import { createUser } from '../api/user';
  
    let username = '';
    let password = '';
    let isCreating = false;
    let msg = ''
    let snackbarWithoutClose;
    const handleCreate = async() => {
      // console.log('22');
      const res = await createUserAjax({username,password})
      if(res.status==200) {
        msg = '创建成功！'
        snackbarWithoutClose.open()
        isCreating = false
        username = ''
        password = ''
      }
      // location.href = '/about'
      
    }
    const handleLogin = async() => {
      try {
        const res = await loginAjax({username,password})
      if(res.status==200) {
        msg = '登录成功！'
        goto(`/person/${res.data.userid}`)
      }
      else msg = res.data
      } catch (error) {
        msg = error.response.data
      }
      snackbarWithoutClose.open()
  
      
    }
    
  </script>
  
  <style lang="scss">
    
    $primary-color: #007bff;
    $primary-color-hover: #0069d9;
    $secondary-color: #6c757d;
    $secondary-color-hover: #5a6268;
  
    body {
      background-color: #f2f2f2;
      font-family: Arial, sans-serif;
    }
  
    .container {
      width: 400px;
      margin: 0 auto;
      padding: 20px;
      background-color: #fff;
      border-radius: 5px;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
      animation: slide-up 0.5s ease-out;
    }
  
    .container h1 {
      text-align: center;
      margin-bottom: 20px;
      color: #333;
    }
  
    .form-group {
      margin-bottom: 20px;
    }
  
    .form-group label {
      display: block;
      margin-bottom: 5px;
      color: #666;
    }
  
    .form-group input {
      width: 380px;
      padding: 10px;
      border: 1px solid #ccc;
      border-radius: 5px;
      font-size: 16px;
      color: #333;
    }
  
    .form-group input:focus {
      outline: none;
      border-color: $primary-color;
      box-shadow: 0 0 5px rgba($primary-color, 0.5);
    }
  
    .btn {
      display: block;
      width: 100%;
      padding: 10px;
      background-color: $primary-color;
      border: none;
      border-radius: 5px;
      font-size: 16px;
      color: #fff;
      cursor: pointer;
      transition: background-color 0.3s ease-out;
    }
  
    .btn:hover {    background-color: $primary-color-hover;
    }
  
    .btn-secondary {
      background-color: $secondary-color;
      border-color: $secondary-color;
      transition: background-color 0.3s ease-out, border-color 0.3s ease-out;
    }
  
    .btn-secondary:hover {
      background-color: $secondary-color-hover;
      border-color: $secondary-color-hover;
    }
  
    @keyframes slide-up {
      from {
        transform: translateY(50px);
        opacity: 0;
      }
      to {
        transform: translateY(0);
        opacity: 1;
      }
    }
  
  </style>
  
  {#if isCreating}
    <div class="container">
      <h1>Create Account</h1>
      <form on:submit={handleCreate}>
        <div class="form-group">
          <label for="username">Username</label>
          <input type="text" id="username" bind:value={username} />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input type="password" id="password" bind:value={password} />
        </div>
        <div class="form-group">
          <button class="btn">
            Create Account
          </button>
        </div>
        <div class="form-group">
          <button type="button" on:click={() => isCreating = false} class="btn btn-secondary">
            Back to Login
          </button>
        </div>
      </form>
    </div>
  {:else}
    <div class="container">
      <h1>Login</h1>
      <form on:submit={handleLogin}>
        <div class="form-group">
          <label for="username">Username</label>
          <input type="text" id="username" bind:value={username} />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input type="password" id="password" bind:value={password} />
        </div>
        <div class="form-group">
          <button class="btn">Login</button>
        </div>
        <div class="form-group">
          <button type="button" on:click={() => isCreating = true} class="btn btn-secondary">
            Create Account
          </button>
        </div>
      </form>
    </div>
  {/if}
  <Snackbar bind:this={snackbarWithoutClose}>
    <Label>{msg}</Label>
  </Snackbar>
  
