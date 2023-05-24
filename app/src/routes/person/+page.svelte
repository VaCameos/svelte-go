<script>
  import { page } from '$app/stores';
    import   Button  from '@smui/button';
    import TextField from '@smui/textfield';
  import { onMount } from 'svelte';
    import Snackbar, {  Label } from '@smui/snackbar';
  import { getUser, updateUserInfo } from '../../util/api';
    let username = ''; // 账号
    let password = ''; // 密码\
    let btnText = '编辑'
    let isEdit = false; 
    let snackbarWithoutClose
    let msg = ''
    const userid = parseInt($page.url.searchParams.get('userid'));
    const getUserInfo = async() => {
      const res = await getUser(userid)
      if(res.status==200){
        let {data} = res
        username = data.username
        password = data.password
      }
    }
    const handleUpdate = async() => {
      if(!isEdit) {
        isEdit = true
        btnText = '保存'
        return
      }
      const res = await updateUserInfo({username,userid,password})
      if(res.status==200) {
        btnText = '编辑'
        isEdit = false
        getUserInfo()
        msg = '保存成功'
        snackbarWithoutClose.open()
      }


    }
    onMount(()=>{
      getUserInfo()
    })
  </script>
  
  <h1>个人信息</h1>
  
  <TextField label="账号" value={username} disabled={!isEdit} on:input={(e) => username = e.target.value}/> <br>
  
  <TextField label="密码" type="password" disabled={!isEdit} value={password} on:input={(e) => password = e.target.value} /> <br>
  
  <Button raised on:click={handleUpdate}>{btnText}</Button>
  <Snackbar bind:this={snackbarWithoutClose}>
    <Label>{msg}</Label>
  </Snackbar>
  