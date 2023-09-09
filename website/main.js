function fetchAndDisplayData() {
    // Получить значение UUID из поля ввода
    const uuid = document.getElementById('uuid').value;
    console.log(uuid)
    console.log(JSON.stringify({uuid: uuid}))

    // Генерировать HTTP-запрос на бэкэнд
    fetch('http://localhost:8080/aboba', {
        method: 'POST',
        body: JSON.stringify({uuid: uuid}),
        headers: {
            'Content-type': 'application/json; charset=UTF-8',
        },
    })
        .then(response => response.json())
        .then(data => {
            // Отобразить полученный JSON на экране
            if (data === null) {
                document.getElementById('result').innerHTML = "Order not found"
            } else {
                str = JSON.stringify(data)
                parsed_obj = JSON.parse(str)


                res =
                    "<h1>Result:<br></h1>" +
                    "Order Id: " + parsed_obj.order_uid + "<br>" +
                    "Entry: " + parsed_obj.entry + "<br>" +
                    "Track number: " + parsed_obj.track_number + "<br>" +
                    "Delivery service: " + parsed_obj.delivery_service + "<br>" +
                    "Date created: " + parsed_obj.date_created + "<br>" +
                    "<h2>Payment info:<br></h2>" +
                    "<ul>" +
                        "<li>" + "Transaction: " + parsed_obj.payment.transaction + "</li>" +
                        "<li>" + "Currency: " + parsed_obj.payment.currency + "</li>" +
                        "<li>" + "Provider: " + parsed_obj.payment.provider + "</li>" +
                        "<li>" + "Amount: " + parsed_obj.payment.amount + "</li>" +
                        "<li>" + "Payment date: " + parsed_obj.payment.payment_dt + "</li>" +
                        "<li>" + "Bank: " + parsed_obj.payment.bank + "</li>" +
                        "<li>" + "Delivery cost: " + parsed_obj.payment.delivery_cost + "</li>" +
                    "</ul>" +
                    "<h2>Delivery info:<br></h2>" +
                    "<ul>" +
                        "<li>" + "Name:  " + parsed_obj.delivery.name + "</li>" +
                        "<li>" + "Phone:  " + parsed_obj.delivery.phone + "</li>" +
                        "<li>" + "Zip:  " + parsed_obj.delivery.zip + "</li>" +
                        "<li>" + "City:  " + parsed_obj.delivery.city + "</li>" +
                        "<li>" + "Address:  " + parsed_obj.delivery.address + "</li>" +
                        "<li>" + "Region:  " + parsed_obj.delivery.region + "</li>" +
                        "<li>" + "Email:  " + parsed_obj.delivery.email + "</li>" +
                    "</ul>" +
                    "<h2>" + "Items: "+ "<br></h2>" + "<ul>"
                ;


                parsed_obj.items.forEach(function (item){
                    res = res +
                        "<li> " + "Item name: " + item.name + "</li>" +
                            "<ul>"+
                                "<li>" + "Brand: " + item.brand + "</li>" +
                                "<li>" + "Size: " + item.size + "</li>" +
                                "<li>" + "Sale: " + item.sale + "</li>" +
                                "<li>" + "Price: " + item.price + "</li>" +
                            "</ul>"
                    ;
                })

                res = res +  "</ul>"


                document.getElementById('result').innerHTML = res

            }


        })
        .catch(error => {
            console.error('Error:', error);
        });
}

