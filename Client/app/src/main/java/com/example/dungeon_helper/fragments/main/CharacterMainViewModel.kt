package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class CharacterMainViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "CHARACTER"
    }
    val text: LiveData<String> = _text
}