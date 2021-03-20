package com.example.demo.controller;

import com.example.demo.model.Restaurant;
import com.example.demo.repository.RestaurantRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/")
public class DemoController {

    @GetMapping("/")
    public String printWelcomeMessage() {
        return "<h1>Hello World!</h1></br><em>This demo site was deployed with <a href=\"https://www.compose-generator.com\" " +
                "target=\"_blank\">Compose Generator</a>!</em></br>Please navigate to <a href=\"restaurant\">/restaurant</a> to see some data.";
    }
}
