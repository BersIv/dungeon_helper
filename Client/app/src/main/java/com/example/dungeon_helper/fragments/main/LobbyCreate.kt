package com.example.dungeon_helper.fragments.main

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
import com.example.dungeon_helper.databinding.FragmentLobbyCreateBinding
import com.example.dungeon_helper.databinding.FragmentLobbyMainBinding

class LobbyCreate : Fragment() {

    companion object {
        fun newInstance() = LobbyCreate()
    }

    private lateinit var viewModel: LobbyCreateViewModel

    private var _binding: FragmentLobbyCreateBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val lobbyMainViewModel = ViewModelProvider(this)[LobbyCreateViewModel::class.java]

        _binding = FragmentLobbyCreateBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textLobbyCreate
        lobbyMainViewModel.text.observe(viewLifecycleOwner) {
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

        val name = binding.textFieldName.editText
        val pwd = binding.textFieldPwd.editText
        val num = binding.textFieldNumber.editText

        val backBtn = binding.backBtn
        val createBtn = binding.createBtn




        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_lobbyCreate_to_navigation_lobby)
        }

    }
}