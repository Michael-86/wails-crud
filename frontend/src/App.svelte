<script>
// Import the router component
import Router from 'svelte-spa-router'


// Import the list of routes
import routes from './routes'

  import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink,
    Dropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem
  } from 'sveltestrap';

	import { Styles } from 'sveltestrap';
	
  let isOpen = false;

  function handleUpdate(event) {
    isOpen = event.detail.isOpen;
  }

</script>

<Styles />
<body>
<Navbar style="background-color: #0275dd" bg-prnd="lg">
  <NavbarBrand href="/">Kabelmärkaren GOLANG</NavbarBrand>
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="lg" on:update={handleUpdate}>
    <Nav class="ms-auto" navbar>
      <NavItem>
        <NavLink href="#components/">Components</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="#/">Home</NavLink>
        <NavLink href="#/hello/svelte">Say hi!</NavLink>
        <NavLink href="#/wild/card">Wildcard route</NavLink>
        <NavLink href="#/does/not/exist">Not found</NavLink>
        <NavLink href="#/Kabel">Kabel</NavLink>
      </NavItem>
      <Dropdown>
        <DropdownToggle nav caret>Options</DropdownToggle>
        <DropdownMenu end>
          <DropdownItem>Option 1 with long text leaking out of the view</DropdownItem>
          <DropdownItem>Option 2</DropdownItem>
          <DropdownItem divider />
          <DropdownItem>Reset</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>

<h1>svelte-spa-router sample</h1>
<h2>Basic routing</h2>
</body>

<!-- Show the router -->
<Router {routes} />

<style>
body {
  background-color: black;
}

h1 {
  color: white;
  text-align: center;
}

h2 {
 color: white;
  text-align: center;
}


</style>
