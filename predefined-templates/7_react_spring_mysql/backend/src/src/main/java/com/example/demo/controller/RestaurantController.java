package com.example.demo.controller;

import com.example.demo.model.Restaurant;
import com.example.demo.repository.RestaurantRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/restaurant")
public class RestaurantController {

    @Autowired
    private RestaurantRepository repository;

    @GetMapping("/")
    public List<Restaurant> getAll() {
        return repository.findAll();
    }

    @PostMapping("/")
    public Restaurant insertOne(@RequestBody Restaurant restaurant) {
        return repository.save(restaurant);
    }
}