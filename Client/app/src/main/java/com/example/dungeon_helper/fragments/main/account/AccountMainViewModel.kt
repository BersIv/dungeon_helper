package com.example.dungeon_helper.fragments.main.account

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AccountMainViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "ACCOUNT"
    }

    val text: LiveData<String> = _text

}