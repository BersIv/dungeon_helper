package com.example.dungeon_helper.fragments.auth

import android.content.Context
import android.content.Intent
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.SharedViewModel
import com.example.dungeon_helper.databinding.FragmentAuthRestorePwdBinding
import com.example.dungeon_helper.databinding.FragmentCharacterMainBinding
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.Interceptor
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException

class AuthRestorePwd : Fragment() {

    companion object {
        fun newInstance() = AuthRestorePwd()
    }

    private lateinit var viewModel: AuthRestorePwdViewModel

    private var _binding: FragmentAuthRestorePwdBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRestorePwdViewModel = ViewModelProvider(this)[AuthRestorePwdViewModel::class.java]

        _binding = FragmentAuthRestorePwdBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textRestorePwd
        authRestorePwdViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }


    override fun onStart() {
        super.onStart()

        val shared = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val mail = binding.textFieldMail.editText
        val backBtn = binding.backBtn
        val restoreBtn = binding.restoreBtn

        backBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd_to_auth)
        }

        restoreBtn.setOnClickListener {
//            GlobalScope.launch(Dispatchers.Main) {
//
//                val client = OkHttpClient()
//
//                val jsonBody = JSONObject().apply {
//                    put("email", mail?.text.toString())
//                }
//
//                val mediaType = "application/json; charset=utf-8".toMediaType()
//                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)
//
//                val request = Request.Builder()
//                    .url("http://194.247.187.44:5000/auth/restorePassword")
//                    .post(body)
//                    .build()
//
//                try {
//                    val response = withContext(Dispatchers.IO) {
//                        client.newCall(request).execute()
//                    }
//
//                    if (!response.isSuccessful) {
//                        throw IOException("Запрос к серверу не был успешен:" +
//                                " ${response.code} ${response.message}")
//                    }
//                    // пример получения конкретного заголовка ответа
//                    println("${response.code} ${response.message}")
//                    // вывод тела ответа
//                    println(response.body!!.string())
//
//                    //(activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd_to_auth)
//                    (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd_to_authRestorePwd2)
//
//                } catch (e: IOException) {
//                    println("Ошибка подключения: $e");
//                }
                shared.mailRestorePwd.value = mail?.text.toString()

                (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd_to_authRestorePwd2)





        }
    }
}