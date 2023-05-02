<!-- https://eugenkiss.github.io/7guis/tasks#crud -->

<script>
    import {GreetPerson, Print1} from "../../wailsjs/go/main/App";

    let people = [
        {
            first: 'Hans',
            last: 'Emil'
        },
        {
            first: 'Max',
            last: 'Mustermann'
        },
        {
            first: 'Roman',
            last: 'Tisch'
        }
    ];


    let prefix = '';
    let first = '';
    let last = '';
    let i = 0;

    $: filteredPeople = prefix
        ? people.filter(person => {
            const name = `${person.last}, ${person.first}`;
            return name.toLowerCase().startsWith(prefix.toLowerCase());
        })
        : people;

    $: selected = filteredPeople[i];

    $: reset_inputs(selected);

    function create() {
        people = people.concat({first, last});
        i = people.length - 1;
        first = last = '';
    }

    function update() {
        selected.first = first;
        selected.last = last;
        people = people;
    }

    function print2() {
        document.getElementById("demo").innerHTML = JSON.stringify(people, null, 4);

        //var a = ["a", "b", "c"];
        //for (i = 0; i < people.length; i++)
        //document.getElementById("demo").innerHTML += JSON.stringify((i+1) + ": " + people[i]);
        //document.writeln((i+1) + ": " + people[i]);
        //document.getElementById("demo").innerHTML = people.first;
    }

	// Passing all people to GO function
    function Print() {
        Print1(people)
    }

    // Passing first item in array to GO function
    function greetPerson() {
        GreetPerson(people[0])
    }

    function remove() {
        // Remove selected person from the source array (people), not the filtered array
        const index = people.indexOf(selected);
        people = [...people.slice(0, index), ...people.slice(index + 1)];

        first = last = '';
        i = Math.min(i, filteredPeople.length - 2);
    }

    function reset_inputs(person) {
        first = person ? person.first : '';
        last = person ? person.last : '';
    }
</script>

<input placeholder="filter prefix" bind:value={prefix}>

<select bind:value={i} size={5}>
    {#each filteredPeople as person, i}
        <option value={i}>{person.first}, {person.last}</option>
    {/each}
</select>
<p>first</p>
<label><input bind:value={first} placeholder="first"></label>
<p>second</p>
<label><input bind:value={last} placeholder="last"></label>

<div class='buttons'>
    <button on:click={create} disabled="{!first || !last}">create</button>
    <button on:click={update} disabled="{!first || !last || !selected}">update</button>
    <button on:click={remove} disabled="{!selected}">delete</button>
    <button class="btn" on:click={print2}>Submit</button>
    <button on:click={Print}>Print</button>
    <button on:click={greetPerson}>Greet first person in array</button>
</div>
<p id="demo"></p>
<style>
    * {
        font-family: inherit;
        font-size: inherit;
    }

    input {
        display: block;
        margin: 0 0 0.5em 0;
    }

    select {
        float: left;
        margin: 0 1em 1em 0;
        width: 14em;
    }

    .buttons {
        clear: both;
    }
</style>
