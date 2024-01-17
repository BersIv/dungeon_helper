package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class CharacterCreate2ViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "Character_create_2"
    }
    val text:LiveData<String> = _text
}