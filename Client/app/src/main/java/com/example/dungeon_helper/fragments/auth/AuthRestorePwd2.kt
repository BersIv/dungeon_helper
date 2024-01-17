package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.SharedViewModel
import com.example.dungeon_helper.databinding.FragmentAuthRestorePwd2Binding
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

class AuthRestorePwd2 : Fragment() {

    companion object {
        fun newInstance() = AuthRestorePwd2()
    }

    private lateinit var viewModel: AuthRestorePwd2ViewModel

    private var _binding: FragmentAuthRestorePwd2Binding? = null
    private  val binding get() = _binding!!
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRestorePwd2ViewModel = ViewModelProvider(this)[AuthRestorePwd2ViewModel::class.java]
        _binding = FragmentAuthRestorePwd2Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textRestorePwd2
        authRestorePwd2ViewModel.text.observe(viewLifecycleOwner)
        {
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

        var mail = ""
        shared.mailRestorePwd.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            mail = it
        })

        val newPwd = binding.textFieldPwd.editText
        val backBtn2 = binding.backBtn2
        val savePwdBtn = binding.savePwdBtn

        backBtn2.setOnClickListener{
            (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd2_to_authRestorePwd)
        }

        savePwdBtn.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {

                val client = OkHttpClient()

                val jsonBody = JSONObject().apply {
                    put("email", mail)
                    put("newPassword", newPwd?.text.toString())
                }

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/auth/restorePassword")
                    .post(body)
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
                    println("${response.code} ${response.message}")
                    // вывод тела ответа
                    println(response.body!!.string())

                    (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd2_to_auth)

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }


            }

        }
    }

}