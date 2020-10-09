<p align="center">
    <a href=""">
        <img height=200 src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQWvuFYj4jQjctg_BrTDkaKWgK0G7Pm3hSGKQ&usqp=CAU">
    </a>
    <br><u>A Fruit store built using Angular 10 for font end and Golang for backend api</u>
</p>



**Overview :** This project is to build an Fruit store with both front end and back end parts of the application. Angular 10 is used to build front end part and Golang is used for building RESTFul API serving data from in-memory DB

**Project Rules**

**Task**

Develop an online eCommerce store selling fruit, which contains the following features:

- Simple sign-up and login form.

- Browse the following products

- - Apples
  - Bananas
  - Pears
  - Oranges

- Add items to your cart

- - Adjust quantity.
  - Delete items from the cart.
  - Apply coupons.

- Checkout your cart

- - Mocked purchase
  - An address does not need to be entered.

**Requirements**

Architecture diagrams.

- Single-page frontend app.
- Backend RESTful web service written in GoLang.
- Users must be able to return to their cart after closing the browser, and see the previous items that were added.

**Cart Rules**

 If 7 or more apples are added to the cart, a 10% discount is applied to all apples.

- For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set.

- - These sets must be added to their own cart item entry.
  - If pears or bananas already exist in the cart, this discount must be recalculated when new pears or bananas are added.

- A coupon code can be used to get a 30% discount on oranges, if applied to the cart, otherwise oranges are full price.

- - Can only be applied once
  - Has an configurable expiry timeout (10 seconds for testing purposes) once generated.

- The following totals must be shown:

- - Total price.
  - Total savings.

**Assumptions**

```shell
1. Node v12.18.2 installed
2. go1.14.4 installed
3. goa design installed https://github.com/goadesign/goa
4. Running Linux distro or MacOS
```

**Build and Run Command**

```sh
chmod +x buildandrun.sh
chmod +x gobuild.sh
./build.sh
```

Now if the build and run is completed successfully application starts up and you would see a screen something like this

![https://res.cloudinary.com/rockey5520/image/upload/v1594848905/fruitstore/successfulbuild_mgcpqc.png]()

![](https://res.cloudinary.com/rockey5520/image/upload/v1594848905/fruitstore/successfulbuild_mgcpqc.png)



To see the application, launch incognito mode in a browser you like and go to `http://localhost:8080`and this should present you a login form as below

![](https://res.cloudinary.com/rockey5520/image/upload/v1594849128/fruitstore/loginform_tzxsg1.jpg)

If you have not created an account earlier you can use the login id you choose (preferred you name without spaces or mobile number) and click login. But if you try to register with an id exists in database form with throw error saying `userid already exists`

Upon login you should see a shopping cart something like below

![](https://res.cloudinary.com/rockey5520/image/upload/v1594849520/fruitstore/home_page_uugebm.jpg)



![](https://res.cloudinary.com/rockey5520/image/upload/v1594851723/fruitstore/discounts_applied_zvonn5.jpg)

Here you can use self explanatory descriptions to add fruits to carts and discounts earlier mentioned will be applied automatically but for 30% discount on oranges , one need to click on `ORANGE 30 Discount coupn`  to apply which is valid for only 10 seconds post the time discount will removed from the cart. 

In Angular Observable is used to link the components so that changes are applied across other components when there is a change to one. Discount coupon table is not in requirement but i left it there so it sits as a nice help to check if the discounts and coupons active for a particular user.



While there is so much more this site can evolve, I am going to see next if i can implement websockets in Golang to inform angular about changes in the backend. If you made it until here , wow!! you had lots of patience , Please feel free to drop me a message if you have any questions