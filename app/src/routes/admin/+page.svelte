

<DataTable table$aria-label="People list" style="max-width: 100%;">
    <Head>
      <Row>
        <Cell>userid</Cell>
        <Cell>username</Cell>
        <Cell numeric>password</Cell>
        <Cell numeric><Button on:click={() => open = true} variant="raised">
          <Label>Add</Label>
          </Button></Cell>
      </Row>
    </Head>
    <Body>
        {#each dataList as item (item.userid)}
          <Row>
            <Cell numeric>{item.userid}</Cell>
            <Cell>{item.username}</Cell>
            <Cell>{item.password}</Cell>
            <Cell>
              <Button on:click={handleDel(item.userid)} variant="raised">
              <Label>Delete</Label>
              </Button>
              <Button on:click={handleUpdate(item)} variant="raised">
              <Label>update</Label>
              </Button>
          </Cell>
          </Row>
        {/each}
      </Body>
  </DataTable>
  <Snackbar bind:this={snackbarWithoutClose}>
    <Label>{msg}</Label>
  </Snackbar>
  <Dialog
  bind:open
  aria-labelledby="default-focus-title"
  aria-describedby="default-focus-content"
>
  <Title id="default-focus-title">{tmpId?'Edit':'Add'}</Title>
  <Content id="default-focus-content">
    <div class="form-group">
      <label for="username">Username</label>
      <input type="text" id="username" bind:value={username} />
    </div>
    <div class="form-group">
      <label for="password">Password</label>
      <input type="password" id="password" bind:value={password} />
    </div>
  </Content>
  <Actions>
    
    <Button
      on:click={handleClose}
    >
      <Label>Close</Label>
    </Button>
    <Button on:click={handleAdd}>
      <Label>Confirm</Label>
    </Button>
  </Actions>
</Dialog>
<style lang="scss">
  
  $primary-color: #007bff;
  $primary-color-hover: #0069d9;
  $secondary-color: #6c757d;
  $secondary-color-hover: #5a6268;

 

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


</style>
  <script lang="js">
    import Dialog, { Title, Content,Actions} from '@smui/dialog';
  import Snackbar from '@smui/snackbar';
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import Button, { Label } from '@smui/button';
    import { onMount } from 'svelte';
  import { createUserAjax, deleteUser, getUserList, updateUserInfo } from '../../util/api';
    let open = false;
    let msg = ''
    let snackbarWithoutClose
    let dataList = [];
    let username = ''
    let password = ''
    let tmpId = ''
    const init = async() => {
      const res = await getUserList()
      if(res.status==200) dataList = res.data
    }
    onMount(()=>{
       init()
    })
    const handleDel = async(id) => {
      const res = await deleteUser(id)
      if(res.status==200) {
        msg = '删除成功'
        snackbarWithoutClose.open()
        init()
      }
    }
    const clearForm = () => {
      username = ''
      password = ''
      tmpId = ''
    }
    const handleClose = () => {
      open = false
      setTimeout(() => {
        clearForm()
        
      }, 100);
    }
    const handleAdd = async() => {
      let res;
      if(tmpId) res = await updateUserInfo({username,password,userid:tmpId})
      else  res = await createUserAjax({username,password})
      if(res.status==200){
         msg = tmpId?'编辑成功':'添加成功'
        snackbarWithoutClose.open()
        clearForm()
        init()
      }
    }
    const handleUpdate = (data) => {
      username = data.username
      password = data.password
      tmpId = data.userid
      open = true
    }
  </script>
  