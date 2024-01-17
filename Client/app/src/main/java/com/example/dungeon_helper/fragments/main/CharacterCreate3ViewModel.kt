package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class CharacterCreate3ViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "Character_create_3"
    }
    val text: LiveData<String> = _text
}