

<DataTable table$aria-label="People list" style="max-width: 100%;">
    <Head>
      <Row>
        <Cell>Name</Cell>
        <Cell>Favorite Color</Cell>
        <Cell numeric>Favorite Number</Cell>
      </Row>
    </Head>
    <Body>
        {#each dataList as item (item.id)}
          <Row>
            <Cell numeric>{item.userid}</Cell>
            <Cell>{item.username}</Cell>
            <Cell>{item.password}</Cell>
            <Cell>
              <Button on:click={() => clicked++} variant="raised">
              <Label>Raised</Label>
              </Button>
          </Cell>
          </Row>
        {/each}
      </Body>
  </DataTable>
  
  <script lang="js">
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import Button, { Label } from '@smui/button';
    let clicked = 0;
    let dataList = []
    for (let i = 0; i < 10; i++) {
      dataList.push({
        userid:i,
        username:`user${i}`,
        password:'123456'
      })
      
    }
  </script>
  