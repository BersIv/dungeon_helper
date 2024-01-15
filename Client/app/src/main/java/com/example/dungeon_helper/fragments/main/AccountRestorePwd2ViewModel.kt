package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AccountRestorePwd2ViewModel : ViewModel() {
    private val _text =  MutableLiveData<String>().apply {
        value = "AccountRestorePwd2"
    }
    val text: LiveData<String> = _text
}