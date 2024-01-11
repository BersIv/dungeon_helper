package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AuthRegistrationViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "AuthRegistration"
    }
    val text: LiveData<String> = _text
}