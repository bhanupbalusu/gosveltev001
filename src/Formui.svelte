<svelte:options tag="my-formui" />

<script>
</script>

<main>
    <header>
        <h1>Enter Product Details</h1>
    </header>
    <div>
        <form action="/action_page.php">
            <label for="pname">Product Name</label>
            <input
                type="text"
                id="pname"
                name="productname"
                placeholder="Name of product ..."
            />
            <label for="pdesc">Product Description</label>
            <textarea
                id="txtArea"
                name="productdesc"
                rows="5"
                cols="70"
                placeholder="Describe about your product ..."
            />
            <fieldset>
                <legend>Quantity Details</legend>
                <div
                    class="row justify-content-start"
                    style="padding-top:10px;"
                >
                    <div class="col">
                        <label for="vol">Bulk Volume</label>
                        <input
                            type="text"
                            id="vol"
                            name="volume"
                            placeholder="Total bulk volume ..."
                        />
                    </div>
                    <div class="col">
                        <label for="unit">Unit</label>
                        <select id="unit" name="unit">
                            <option value="kg">Kg</option>
                            <option value="l">Litres</option>
                            <option value="pcs">Pieces</option>
                        </select>
                    </div>
                </div>
            </fieldset>
            <fieldset>
                <legend>Price Details</legend>
                <div
                    class="row justify-content-start"
                    style="padding-top:10px;"
                >
                    <div class="col" style="padding-top:10px;">
                        <label for="price">Price Value</label>
                        <input
                            type="text"
                            id="price"
                            name="pricevalue"
                            placeholder="Price charged ..."
                        />
                    </div>
                    <div class="col" style="padding-top:10px;">
                        <label for="curr">Currency</label>
                        <select id="curr" name="currency">
                            <option value="gbp">GBP</option>
                            <option value="usd">USD</option>
                            <option value="eur">EUR</option>
                            <option value="inr">INR</option>
                        </select>
                    </div>
                </div>
                <div
                    class="row justify-content-start"
                    style="padding-top:10px;"
                >
                    <div class="col" style="padding-top:10px;">
                        <label for="perunit">Base Quantity Value Per Unit</label
                        >
                        <input
                            type="text"
                            id="perunit"
                            name="qtyperunit"
                            placeholder="Minimum base quantity that is sold ..."
                        />
                    </div>
                    <div class="col" style="padding-top:10px;">
                        <label for="unit">Unit</label>
                        <select id="unit" name="unit">
                            <option value="kg">Kg</option>
                            <option value="gm">Grams</option>
                            <option value="l">Litres</option>
                            <option value="ml">Milli Litres</option>
                            <option value="pcs">Piece(s)</option>
                        </select>
                    </div>
                </div>
            </fieldset>
            <div class="submit">
                <input type="submit" value="Submit" />
            </div>
        </form>
    </div>
    <footer>
        <h1>This is the Footer!!!!!!</h1>
    </footer>
</main>

<style>
    main {
        background-color: rgb(169, 184, 201);
    }
    input[type="text"],
    textarea,
    select {
        width: 100%;
        padding: 12px 20px;
        border: 1px solid #ccc;
        margin: 8px 0;
        display: inline-block;
        border-radius: 5px;
        box-sizing: border-box;
    }

    input[type="submit"] {
        width: 20%;
        background-color: #4caf50;
        color: white;
        padding: 15px 10px;
        margin-bottom: 10px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        float: right;
    }

    input[type="submit"]:hover {
        background-color: #45a049;
    }

    .submit {
        padding: auto;
        margin: 20px;
    }

    fieldset {
        padding: 15px;
        margin: auto;
        margin-top: 10px;
        border-radius: 5px;
        border: 1px solid rgb(222, 181, 247);
    }

    div {
        border-radius: 5px;
        background-color: #f2f2f2;
        padding: 20px;
    }

    h1 {
        margin: auto;
        padding: 10px;
        text-align: center;
        color: rgb(31, 31, 184);
    }
</style>
