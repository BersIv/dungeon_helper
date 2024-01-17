package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.SharedViewModel
import com.example.dungeon_helper.databinding.FragmentAccountEditBinding
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException

class AccountEdit : Fragment() {

    companion object {
        fun newInstance() = AccountEdit()
    }

    private lateinit var viewModel: AccountEditViewModel

    private var _binding: FragmentAccountEditBinding? = null
    private val binding get() = _binding!!
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isVisible = false
        menuItem2.isVisible = false
        menuItem3.isVisible = false
        menuItem4.isVisible = false
       val accountEditViewModel = ViewModelProvider(this)[AccountEditViewModel::class.java]
        _binding = FragmentAccountEditBinding.inflate(inflater, container, false)
        val root: View =  binding.root
        val textView: TextView = binding.textAccountEdit
        accountEditViewModel.text.observe(viewLifecycleOwner)
        {
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isVisible = true
        menuItem2.isVisible = true
        menuItem3.isVisible = true
        menuItem4.isVisible = true
    }

    override fun onStart() {
        super.onStart()

        val shared = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val exAccBtn = binding.exAccBtn
        val exEditBtn = binding.exEditBtn
        val changePwdBtn = binding.changePwdBtn
        val newNick = binding.textFieldNick.editText

        exAccBtn.setEnabled(false)
        changePwdBtn.setEnabled(false)

        exEditBtn.setOnClickListener{
            GlobalScope.launch(Dispatchers.Main) {

                //val client = OkHttpClient()
                val jsonBody = JSONObject().apply {
                    put("nickname", newNick?.text.toString())
                }

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                var token: String = ""
                shared.token.observe(viewLifecycleOwner, Observer {
                    // updating data in displayMsg
                    token = it
                })


                //println(getCookieValue())

                val client = OkHttpClient.Builder().addInterceptor { chain ->
                    val original = chain.request()
                    val authorized = original.newBuilder()
                        .header("Authorization", token)
                        .build()
                    chain.proceed(authorized)
                }
                    .build()

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/account/change/nickname")
                    .patch(body)
                    .build()

                try {
                    val response = withContext(Dispatchers.IO) {
                        client.newCall(request).execute()
                    }

                    if (!response.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response.code} ${response.message}")
                    }
                    // пример получения конкретного заголовка ответа
                    println("Response message: ${response.header("message")}")
                    // вывод тела ответа
                    println(response.body!!.string())
                    shared.nickname.value = newNick?.text.toString()

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }
            (activity as MainActivity).navController.navigate(R.id.action_accountEdit_to_navigation_account)
        }

    }

}